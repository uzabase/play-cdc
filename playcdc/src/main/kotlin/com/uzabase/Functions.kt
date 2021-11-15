package com.uzabase

import com.github.tomakehurst.wiremock.client.MappingBuilder
import kotlin.io.path.Path
import kotlin.io.path.createDirectory

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
    return RequestJson(mappingBuilder.build().request.url)
}

interface Writer {
    fun createDirectory()
    fun write(requestJson: RequestJson)
}
