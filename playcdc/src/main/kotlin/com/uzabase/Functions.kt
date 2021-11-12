package com.uzabase

import com.github.tomakehurst.wiremock.client.MappingBuilder
import kotlin.io.path.Path
import kotlin.io.path.createDirectory

fun storeMock(mappingBuilder: MappingBuilder) {
    val folderName = getFolderName()
    val path = folderName?.let { Path(BASE_PATH).resolve(it) } ?: return
    path.createDirectory()

    storeMock(mappingBuilder, FileWriter(path))
}

internal fun storeMock(mappingBuilder: MappingBuilder, writer: Writer) {
    val url = mappingBuilder.build().request.url
    writer.writeRequestPath(url)
}

interface Writer {
    fun writeRequestPath(requestPath: String)
}
