package com.uzabase.playcdc.internal.infra

import com.uzabase.playcdc.internal.Contract

fun verifyResponse(response: Contract.Response, status: Int, body: String?, headers: Map<String, String>?) {
    if (response.status != status) {
        throw AssertionError(
            "Statues were not equal.\n" +
                    "Expected: ${response.status}\n" +
                    "  Actual: $status"
        )
    }

    if (!actualBodyContainsContractsBody(response, body?.ifEmpty { null })) {
        throw AssertionError(
            "Actual body doesn't contain expected body.\n" +
                    "Expected: ${response.jsonBody}\n" +
                    "  Actual: ${body?.let(::toMap)}"
        )
    }

    if (!actualHeadersContainContractsHeaders(response, headers)) {
        throw AssertionError(
            "Actual headers doesn't contain expected headers.\n" +
                    "Expected: ${response.headers}\n" +
                    "  Actual: $headers"
        )
    }
}

private fun actualBodyContainsContractsBody(response: Contract.Response, body: String?) =
    response.jsonBody == null && body == null ||
            (response.jsonBody != null && body != null &&
                    actualBodyContainsContractsBody(response.jsonBody, body.let(::toMap)))

private fun actualBodyContainsContractsBody(contractBody: Map<*, *>, body: Map<*, *>) =
    contractBody.entries.all { body.contains(it) }

private fun <T> Map<*, *>.contains(entry: Map.Entry<*, T>): Boolean {
    val actualValue = this[entry.key]
    val contractValue = entry.value
    return when (actualValue) {
        is Map<*, *> -> contractValue is Map<*, *> && actualBodyContainsContractsBody(contractValue, actualValue)
        else -> actualValue == contractValue
    }
}

private fun actualHeadersContainContractsHeaders(response: Contract.Response, headers: Map<String, String>?) =
    (response.headers == null) ||
            (headers != null && actualHeadersContainContractsHeaders(headers, response.headers))

private fun actualHeadersContainContractsHeaders(actual: Map<String, String>, expected: Map<String, String>) =
    expected.entries.all { (expectedKey, expectedValue) ->
        actual.entries.any { (actualKey, actualValue) ->
            actualKey.equals(expectedKey, ignoreCase = true) && actualValue == expectedValue
        }
    }
