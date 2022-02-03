package com.uzabase.playcdc.internal.infra

import com.uzabase.playcdc.internal.Contract

fun verifyResponse(response: Contract.Response, status: Int, body: String?, headers: Map<String, String>?) {
    if (response.status != status) {
        throw AssertionError(
            "Statues were not equal.\n" +
                    "Expected: ${response.status}\n" +
                    "  Actual: $status")
    }

    if (response.jsonBody != body?.let(::toMap)) {
        throw AssertionError(
            "Bodies were not equal.\n" +
                    "Expected: ${response.jsonBody}\n" +
                    "  Actual: ${body?.let(::toMap)}")
    }

    if (response.headers != headers) {
        throw AssertionError(
            "Headers were not equal.\n" +
                "Expected: ${response.headers}\n" +
                "  Actual: $headers")
    }
}
