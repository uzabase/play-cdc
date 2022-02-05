package com.uzabase.playcdc

import com.uzabase.playcdc.internal.Contract
import com.uzabase.playcdc.internal.infra.sendRequest
import com.uzabase.playcdc.internal.infra.toObject
import com.uzabase.playcdc.internal.infra.verifyResponse

data class Response(
    val status: Int,
    val body: String?,
    val headers: Map<String, String>
)

object PlayCdc {
    fun sendRequest(endpoint: String, contractJson: String, body: String? = null): Response {
        return sendRequest(endpoint, toObject(contractJson, Contract::class.java).request, body)
    }

    fun verifyResponse(contractJson: String, status: Int, body: String? = null, headers: Map<String, String>? = null) {
        verifyResponse(toObject(contractJson, Contract::class.java).response, status, body, headers)
    }
}
