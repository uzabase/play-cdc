package com.uzabase.playcdc

import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.uzabase.playcdc.internal.FileWriter
import com.uzabase.playcdc.internal.getFolderName
import kotlin.io.path.Path


object PlayCdc {
    private const val BASE_PATH = "/tmp"

    fun storeMock(mappingBuilder: MappingBuilder) {
        val folderName = getFolderName()
        val path = folderName?.let { Path(BASE_PATH).resolve(it) } ?: return

        com.uzabase.playcdc.internal.storeMock(mappingBuilder, FileWriter(path))
    }
}
