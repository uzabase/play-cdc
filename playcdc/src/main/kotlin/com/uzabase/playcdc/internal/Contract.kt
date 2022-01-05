package com.uzabase.playcdc.internal

data class Contract(
    val request: Request,
    val response: Response
)

data class Request(
    val url: String,
    val method: String
)

data class Response(
    val status: Int,
    val headers: Map<String, String>,
    val jsonBody: Map<String, Any>
)
