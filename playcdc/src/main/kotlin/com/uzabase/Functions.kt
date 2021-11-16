package com.uzabase

import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.github.tomakehurst.wiremock.matching.RequestPattern
import kotlin.io.path.Path

fun storeMock(mappingBuilder: MappingBuilder) {
    val folderName = getFolderName()
    val path = folderName?.let { Path(BASE_PATH).resolve(it) } ?: return

    storeMock(mappingBuilder, FileWriter(path))
}

internal fun storeMock(mappingBuilder: MappingBuilder, writer: Writer) {
    writer.createDirectory()

    val json = toRequestJson(mappingBuilder)
    writer.write(json)
}

fun toRequestJson(mappingBuilder: MappingBuilder): RequestJson {
    return mappingBuilder.build().request.let {
        RequestJson(
            "${it.url}${queryParameters(it)}",
            it.method.value(),
            headers(it),
            body(it)
        )
    }
}

private fun queryParameters(requestPattern: RequestPattern) =
    requestPattern.queryParameters?.run {
        if (isNotEmpty()) "?" + map { "${it.key}=${it.value.valuePattern.value}" }.joinToString("&")
        else ""
    } ?: ""

private fun headers(requestPattern: RequestPattern) =
    requestPattern.headers?.map { it.key to it.value.valuePattern.value }?.toMap() ?: emptyMap()

private fun body(requestPattern: RequestPattern): String? =
    requestPattern.bodyPatterns?.map { it.value as String }?.firstOrNull()

interface Writer {
    fun createDirectory()
    fun write(requestJson: RequestJson)
}
