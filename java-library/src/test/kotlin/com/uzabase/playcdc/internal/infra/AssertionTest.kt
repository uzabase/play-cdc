package com.uzabase.playcdc.internal.infra

import com.uzabase.playcdc.internal.Contract
import io.kotest.assertions.throwables.shouldThrow
import io.kotest.core.spec.style.FreeSpec

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

                shouldThrow<AssertionError> { verifyResponse(response, 201, null, null) }
            }
        }

        "with body" - {
            "passes when bodies are equal" {
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

            "passes when actual body contains contract's body" - {
                "single level" {
                    val response = Contract.Response(
                        200,
                        null,
                        mapOf(
                            "key1" to "value1"
                        )
                    )

                    val body = """
                        {
                            "key1": "value1",
                            "key2": "value2"
                        }
                    """.trimIndent()

                    verifyResponse(response, 200, body, null)
                }

                "multi level" {
                    val response = Contract.Response(
                        200,
                        null,
                        mapOf(
                            "key1" to mapOf(
                                "key2" to "value2"
                            )
                        )
                    )

                    val body = """
                        {
                            "key1": {
                              "key2": "value2",
                              "key3": "value3"
                            }
                        }
                    """.trimIndent()

                    verifyResponse(response, 200, body, null)
                }
            }

            "passes when contract's body is null and actual body is empty" {
                val response = Contract.Response(
                    200,
                    null,
                    null
                )

                val body = ""

                verifyResponse(response, 200, body, null)
            }

            "fails when bodies are not equal" - {
                val response = Contract.Response(
                    200,
                    null,
                    mapOf(
                        "key" to "value"
                    )
                )

                val body = """
                    {
                        "key": "another value"
                    }
                """.trimIndent()

                shouldThrow<AssertionError> { verifyResponse(response, 200, body, null) }
            }

            "fails when keys match but contents unmatch " - {
                val response = Contract.Response(
                    200,
                    null,
                    mapOf(
                        "key" to "value"
                    )
                )

                val body = """
                    {
                        "key": {
                          "anotherKey": "value"
                        }
                    }
                """.trimIndent()

                shouldThrow<AssertionError> { verifyResponse(response, 200, body, null) }
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

            "passes even when case of key is different between actual and contract's" {
                val response = Contract.Response(
                    200,
                    mapOf(
                        "KEY1" to "value1"
                    ),
                    null
                )

                verifyResponse(response, 200, null, mapOf(
                    "key1" to "value1",
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

                shouldThrow<AssertionError> {
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

                shouldThrow<AssertionError> { verifyResponse(response, 200, null, null) }
            }
        }
    }
})
