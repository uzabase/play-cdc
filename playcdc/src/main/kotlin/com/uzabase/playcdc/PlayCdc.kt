package com.uzabase.playcdc

import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.uzabase.playcdc.internal.Contract
import com.uzabase.playcdc.internal.FileWriter
import com.uzabase.playcdc.internal.legacy.Request
import com.uzabase.playcdc.internal.Writer
import com.uzabase.playcdc.internal.infra.*
import kotlin.io.path.Path
import com.uzabase.playcdc.internal.legacy.Response as InternalResponse

data class Response(
    val status: Int,
    val body: String?
)

object PlayCdc {
    fun storeMock(mappingBuilder: MappingBuilder) {
        storeMock(mappingBuilder, null)
    }

    fun storeMock(mappingBuilder: MappingBuilder, folderName: String?) {
        fileWriter(folderName)?.let {
            it.setup()
            it.write(mappingBuilder.toRequest())
            it.write(mappingBuilder.toResponse())
        }
    }

    fun legacySendRequest(endpoint: String, requestJson: String): Response {
        return sendRequest(endpoint, toObject(requestJson, Request::class.java))
    }

    fun sendRequest(endpoint: String, contractJson: String): Response {
        return sendRequest(endpoint, toObject(contractJson, Contract::class.java).request)
    }

    fun verifyResponse(responseJson: String, status: Int, body: String? = null) {
        verifyResponse(toObject(responseJson, InternalResponse::class.java), status, body)
    }

    private fun fileWriter(folderName: String? = null): Writer? = (folderName ?: findFolderName())
        ?.let { Path(getBasePath()).resolve(it) }
        ?.let(::FileWriter)

    private fun getBasePath(): String = System.getenv("PLAY_CDC_BASE_PATH") ?: "/tmp"
}
