package com.uzabase.playcdc.internal.infra

import com.uzabase.playcdc.internal.Contract
import io.kotest.core.spec.style.FreeSpec
import org.junit.jupiter.api.assertThrows

class AssertionTest : FreeSpec({
    "verify response" - {
        "status" - {
            "passes" {
                val response = Contract.Response(
                    200,
                    null,
                    null
                )

                verifyResponse(response, 200, null, null)
            }

            "fails" {
                val response = Contract.Response(
                    200,
                    null,
                    null
                )

                assertThrows<AssertionError> { verifyResponse(response, 201, null, null) }
            }
        }

        "with body" - {
            "passes" {
                val response = Contract.Response(
                    200,
                    null,
                    mapOf(
                        "mapKey" to mapOf(
                            "key" to "value"
                        ),
                        "arrayKey" to listOf(
                            mapOf("key1" to "value1", "key2" to "value2"),
                            mapOf("key1" to "value21", "key2" to "value22")
                        )
                    )
                )

                val body = """
                    {
                        "mapKey": {
                            "key": "value"
                        },
                        "arrayKey": [
                            {
                                "key1": "value1",
                                "key2": "value2"
                            },
                            {
                                "key1": "value21",
                                "key2": "value22"
                            }
                        ]
                    }
                """.trimIndent()

                verifyResponse(response, 200, body, null)
            }

            "fails" - {
                val response = Contract.Response(
                    200,
                    null,
                    mapOf(
                        "key" to "value"
                    )
                )

                val body = """
                    {
                        "key":  "another value"
                    }
                """.trimIndent()

                assertThrows<AssertionError> { verifyResponse(response, 200, body, emptyMap()) }
            }
        }

        "with headers" - {
            "passes when actual headers contains contract's headers" {
                val response = Contract.Response(
                    200,
                    mapOf(
                        "key1" to "value1"
                    ),
                    null
                )

                verifyResponse(response, 200, null, mapOf(
                    "key1" to "value1",
                    "key2" to "value2"
                ))
            }

            "passes when contract's headers is null" {
                val response = Contract.Response(
                    200,
                    null,
                    null
                )

                verifyResponse(response, 200, null, mapOf(
                    "key1" to "value1",
                    "key2" to "value2"
                ))
            }

            "fails when actual headers doesn't contain contract's headers" {
                val response = Contract.Response(
                    200,
                    mapOf(
                        "key" to "value"
                    ),
                    null
                )

                assertThrows<AssertionError> {
                    verifyResponse(response, 200, null, mapOf(
                        "key" to "another value"
                    ))
                }
            }

            "fails when contract's headers is not null and actual headers is null" {
                val response = Contract.Response(
                    200,
                    mapOf(
                        "key1" to "value1"
                    ),
                    null
                )

                assertThrows<AssertionError> {
                    verifyResponse(response, 200, null, null)
                }
            }
        }
    }
})
