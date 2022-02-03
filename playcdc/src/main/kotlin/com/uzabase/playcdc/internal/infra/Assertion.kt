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

    if (!actualHeadersContainContractsHeaders(response, headers)) {
        throw AssertionError(
            "Actual headers doesn't contain expected headers.\n" +
                    "Expected: ${response.headers}\n" +
                    "  Actual: $headers"
        )
    }
}

private fun actualHeadersContainContractsHeaders(response: Contract.Response, headers: Map<String, String>?) =
    (response.headers == null) || (headers != null && headers.entries.containsAll(response.headers.entries))
