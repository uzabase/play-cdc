package com.uzabase.playcdc.internal.infra

import com.uzabase.playcdc.Response
import com.uzabase.playcdc.internal.Request
import okhttp3.OkHttpClient

private val CLIENT = OkHttpClient.Builder().build()

fun sendRequest(endpoint: String, request: Request): Response {
    return toOkHttp3Request(endpoint, request)
        .let { CLIENT.newCall(it).execute() }
        .let { Response(it.code, it.body?.string()) }
}

private fun toOkHttp3Request(endpoint: String, request: Request) = okhttp3.Request.Builder()
    .url(endpoint + request.url)
    .method(request.method, null)
    .build()
