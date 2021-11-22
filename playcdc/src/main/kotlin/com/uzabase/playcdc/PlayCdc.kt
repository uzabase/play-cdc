package com.uzabase.playcdc

import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.uzabase.playcdc.internal.FileWriter
import com.uzabase.playcdc.internal.getBasePath
import com.uzabase.playcdc.internal.getFolderName
import kotlin.io.path.Path

object PlayCdc {
    fun storeMock(mappingBuilder: MappingBuilder) {
        val basePath = getBasePath()
        val folderName = getFolderName()
        val path = folderName?.let { Path(basePath).resolve(it) } ?: return

        com.uzabase.playcdc.internal.storeMock(mappingBuilder, FileWriter(path))
    }
}
