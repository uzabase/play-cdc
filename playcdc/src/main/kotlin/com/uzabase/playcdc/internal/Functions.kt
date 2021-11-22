package com.uzabase.playcdc.internal

import com.fasterxml.jackson.databind.ObjectMapper
import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.github.tomakehurst.wiremock.http.ResponseDefinition
import com.github.tomakehurst.wiremock.matching.RequestPattern

internal fun storeMock(mappingBuilder: MappingBuilder, writer: Writer) {
    writer.setup()

    val request = toRequestJson(mappingBuilder)
    writer.write(request)

    val response = toResponseJson(mappingBuilder)
    writer.write(response)
}

fun toRequestJson(mappingBuilder: MappingBuilder): RequestJson {
    return mappingBuilder.build().request.let {
        RequestJson(
            "${it.url}${it.toQueryParameters()}",
            it.method.value(),
            it.toHeaders(),
            it.toBody()
        )
    }
}

fun toResponseJson(mappingBuilder: MappingBuilder): ResponseJson {
    return mappingBuilder.build().response.let {
        ResponseJson(
            it.toHeaders(),
            it.toBody()
        )
    }
}

private fun ResponseDefinition.toHeaders() = headers?.all()?.associate { it.key() to it.values().first() } ?: emptyMap()
private fun ResponseDefinition.toBody() = body?.let(::toMap) ?: emptyMap()

private fun RequestPattern.toQueryParameters() =
    queryParameters?.run {
        if (isNotEmpty()) "?" + map { "${it.key}=${it.value.valuePattern.value}" }.joinToString("&")
        else ""
    } ?: ""

private fun RequestPattern.toHeaders() = headers?.map { it.key to it.value.valuePattern.value }?.toMap() ?: emptyMap()

private fun RequestPattern.toBody(): Map<String,Any> = bodyPatterns
    ?.map { it.value  as String }
    ?.map(::toMap)
    ?.firstOrNull() ?: emptyMap()

@Suppress("UNCHECKED_CAST")
private fun toMap(jsonString: String) = ObjectMapper().readValue(jsonString, Map::class.java) as Map<String, Any>

interface Writer {
    fun setup()
    fun write(requestJson: RequestJson)
    fun write(responseJson: ResponseJson)
}
