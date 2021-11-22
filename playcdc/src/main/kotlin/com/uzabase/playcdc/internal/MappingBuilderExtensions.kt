package com.uzabase.playcdc.internal

import com.fasterxml.jackson.databind.ObjectMapper
import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.github.tomakehurst.wiremock.http.ResponseDefinition
import com.github.tomakehurst.wiremock.matching.RequestPattern

fun MappingBuilder.toRequestJson(): RequestJson = build()
    .request.let {
        RequestJson(
            "${it.url}${it.toQueryParameters()}",
            it.method.value(),
            it.toHeaders(),
            it.toBody()
        )
    }

fun MappingBuilder.toResponseJson(): ResponseJson = build()
    .response.let {
        ResponseJson(
            it.toHeaders(),
            it.toBody()
        )
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

private val MAPPER = ObjectMapper()

@Suppress("UNCHECKED_CAST")
private fun toMap(jsonString: String) = MAPPER.readValue(jsonString, Map::class.java) as Map<String, Any>
