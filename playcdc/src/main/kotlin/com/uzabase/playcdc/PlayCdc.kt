package com.uzabase.playcdc

import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.uzabase.playcdc.internal.*
import kotlin.io.path.Path

object PlayCdc {
    fun storeMock(mappingBuilder: MappingBuilder) {
        fileWriter()?.let {
            it.setup()
            it.write(mappingBuilder.toRequestJson())
            it.write(mappingBuilder.toResponseJson())
        }
    }

    private fun fileWriter(): Writer? = findFolderName()
        ?.let { Path(getBasePath()).resolve(it) }
        ?.let(::FileWriter)

    private fun getBasePath(): String = System.getenv("PLAY_CDC_BASE_PATH") ?: "/tmp"
}
