package com.uzabase

data class RequestJson(
    val url: String,
    val method: String,
    val headers: Map<String, String>,
    val body: Map<String, Any>
)
