package com.uzabase.playcdc.internal.infra

import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.github.tomakehurst.wiremock.http.ResponseDefinition
import com.github.tomakehurst.wiremock.matching.RequestPattern
import com.uzabase.playcdc.internal.legacy.Request
import com.uzabase.playcdc.internal.legacy.Response

fun MappingBuilder.toRequest(): Request = build()
    .request.let {
        Request(
            "${it.url}${it.toQueryParameters()}",
            it.method.value(),
            it.toHeaders(),
            it.toBody()
        )
    }

fun MappingBuilder.toResponse(): Response = build()
    .response.let {
        Response(
            it.status,
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
