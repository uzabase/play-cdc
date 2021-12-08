package com.uzabase.playcdc

import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.uzabase.playcdc.internal.FileWriter
import com.uzabase.playcdc.internal.Request
import com.uzabase.playcdc.internal.Writer
import com.uzabase.playcdc.internal.infra.*
import kotlin.io.path.Path

object PlayCdc {
    fun storeMock(mappingBuilder: MappingBuilder) {
        fileWriter()?.let {
            it.setup()
            it.write(mappingBuilder.toRequest())
            it.write(mappingBuilder.toResponse())
        }
    }

    fun sendRequest(endpoint: String, requestJson: String) {
        sendRequest(endpoint, toObject(requestJson, Request::class.java))
    }

    private fun fileWriter(): Writer? = findFolderName()
        ?.let { Path(getBasePath()).resolve(it) }
        ?.let(::FileWriter)

    private fun getBasePath(): String = System.getenv("PLAY_CDC_BASE_PATH") ?: "/tmp"
}
