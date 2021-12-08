package com.uzabase.playcdc

import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.uzabase.playcdc.internal.*
import com.uzabase.playcdc.internal.infra.toObject
import okhttp3.OkHttpClient
import kotlin.io.path.Path
import okhttp3.Request as OkHttp3Request

object PlayCdc {
    fun storeMock(mappingBuilder: MappingBuilder) {
        fileWriter()?.let {
            it.setup()
            it.write(mappingBuilder.toRequest())
            it.write(mappingBuilder.toResponse())
        }
    }

    fun sendRequest(requestJson: String) {
        toObject(requestJson, Request::class.java)
            .let(::toOkHttp3Request)
            .let { CLIENT.newCall(it).execute() }
    }

    private val CLIENT = OkHttpClient.Builder().build()
    private const val ENDPOINT = "http://localhost:8080"

    private fun toOkHttp3Request(request: Request) = OkHttp3Request.Builder()
        .url(ENDPOINT + request.url)
        .method(request.method, null)
        .build()

    private fun fileWriter(): Writer? = findFolderName()
        ?.let { Path(getBasePath()).resolve(it) }
        ?.let(::FileWriter)

    private fun getBasePath(): String = System.getenv("PLAY_CDC_BASE_PATH") ?: "/tmp"
}
