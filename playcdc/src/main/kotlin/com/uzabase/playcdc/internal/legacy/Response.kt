package com.uzabase.playcdc.internal.legacy

data class Response(
    val status: Int,
    val headers: Map<String, String>,
    val body: Map<String, Any>
)
