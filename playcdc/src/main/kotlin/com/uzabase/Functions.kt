package com.uzabase

import com.fasterxml.jackson.databind.ObjectMapper
import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.github.tomakehurst.wiremock.http.ResponseDefinition
import com.github.tomakehurst.wiremock.matching.RequestPattern
import kotlin.io.path.Path

fun storeMock(mappingBuilder: MappingBuilder) {
    val folderName = getFolderName()
    val path = folderName?.let { Path(BASE_PATH).resolve(it) } ?: return

    storeMock(mappingBuilder, FileWriter(path))
}

internal fun storeMock(mappingBuilder: MappingBuilder, writer: Writer) {
    writer.createDirectory()

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
            it.toHeaders()
        )
    }
}

private fun ResponseDefinition.toHeaders() = headers?.all()?.associate { it.key() to it.values().first() } ?: emptyMap()

private fun RequestPattern.toQueryParameters() =
    queryParameters?.run {
        if (isNotEmpty()) "?" + map { "${it.key}=${it.value.valuePattern.value}" }.joinToString("&")
        else ""
    } ?: ""

private fun RequestPattern.toHeaders() = headers?.map { it.key to it.value.valuePattern.value }?.toMap() ?: emptyMap()

private fun RequestPattern.toBody(): Map<String,Any> = bodyPatterns
    ?.map { it.value  as String }
    ?.map { ObjectMapper().readValue(it, Map::class.java) as Map<String, Any> }
    ?.firstOrNull() ?: emptyMap()

interface Writer {
    fun createDirectory()
    fun write(requestJson: RequestJson)
    fun write(requestJson: ResponseJson)
}
