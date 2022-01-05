package com.uzabase.playcdc.internal.infra

import com.uzabase.playcdc.Response
import com.uzabase.playcdc.internal.Request
import com.uzabase.playcdc.internal.legacy.Request as LegacyRequest
import okhttp3.Headers.Companion.toHeaders
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.OkHttpClient
import okhttp3.RequestBody.Companion.toRequestBody

private val CLIENT = OkHttpClient.Builder().build()

fun sendRequest(endpoint: String, request: LegacyRequest): Response {
    return toOkHttp3Request(endpoint, request)
        .let { CLIENT.newCall(it).execute() }
        .let { Response(it.code, it.body?.string()) }
}

fun sendRequest(endpoint: String, request: Request): Response {
    return toOkHttp3Request(endpoint, request)
        .let { CLIENT.newCall(it).execute() }
        .let { Response(it.code, it.body?.string()) }
}

private fun toOkHttp3Request(endpoint: String, request: Request) = okhttp3.Request.Builder()
    .url(endpoint + request.url)
    .method(request.method, null)
    .build()

private fun toOkHttp3Request(endpoint: String, request: LegacyRequest) = okhttp3.Request.Builder()
    .url(endpoint + request.url)
    .method(request.method, request.body?.toJsonString()?.toRequestBody(request.contentType?.toMediaType()))
    .headers(request.headers.toHeaders())
    .build()
