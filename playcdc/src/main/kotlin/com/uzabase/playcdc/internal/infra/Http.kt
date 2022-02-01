package com.uzabase.playcdc.internal.infra

import com.uzabase.playcdc.Response
import com.uzabase.playcdc.internal.Request
import okhttp3.OkHttpClient
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.internal.EMPTY_REQUEST

private val CLIENT = OkHttpClient.Builder().build()

fun sendRequest(endpoint: String, request: Request, body: String?): Response {
    return toOkHttp3Request(endpoint, request, body)
        .let { CLIENT.newCall(it).execute() }
        .let { Response(it.code, it.body?.string()) }
}

private fun toOkHttp3Request(endpoint: String, request: Request, body: String?) = okhttp3.Request.Builder()
    .url(endpoint + request.url)
    .method(request, body)
    .build()

private fun okhttp3.Request.Builder.method(request: Request, body: String?) =
    method(request.method, requestBody(body, request.method))

private fun requestBody(body: String?, method: String) =
    body?.toRequestBody() ?: when (method) {
        "GET" -> null
        else -> EMPTY_REQUEST
    }
