package com.uzabase.playcdc.internal.infra

import com.uzabase.playcdc.Response
import com.uzabase.playcdc.internal.Contract
import java.net.URI
import java.net.http.HttpClient
import java.net.http.HttpHeaders
import java.net.http.HttpRequest
import java.net.http.HttpResponse

private val CLIENT = HttpClient.newHttpClient()

fun sendRequest(endpoint: String, request: Contract.Request, body: String?): Response {
    return toHttpRequest(endpoint, request, body)
        .let { CLIENT.send(it, HttpResponse.BodyHandlers.ofString()) }
        .let { Response(it.statusCode(), it.body(), it.headers().toMap()) }
}

private fun toHttpRequest(endpoint: String, request: Contract.Request, body: String?) =
    HttpRequest.newBuilder(URI.create(endpoint + request.url))
            .method(request, body)
            .headers(request)
            .build()

private fun HttpRequest.Builder.method(request: Contract.Request, body: String?) =
    method(request.method, bodyPublisher(body, request.method))

private fun bodyPublisher(body: String?, method: String): HttpRequest.BodyPublisher =
    when (method) {
        "GET" -> HttpRequest.BodyPublishers.noBody()
        else -> body
            ?.let { HttpRequest.BodyPublishers.ofString(it) }
            ?: HttpRequest.BodyPublishers.noBody()
    }

private fun HttpRequest.Builder.headers(request: Contract.Request) =
    if (request.headers == null) this
    else headers(*toArray(request.headers))

internal fun toArray(headers: Map<String, Map<String, String>>) = headers
    .entries
    .map { (key, value) -> listOf(key, value.values.first()) }
    .flatten()
    .toTypedArray()

private fun HttpHeaders.toMap() = map().map { (key, value) -> key to value.first() }.toMap()
