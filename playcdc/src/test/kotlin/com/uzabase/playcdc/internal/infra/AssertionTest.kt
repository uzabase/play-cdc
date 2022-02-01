package com.uzabase.playcdc.internal.infra

import com.uzabase.playcdc.internal.Response
import io.kotest.core.spec.style.FreeSpec
import org.junit.jupiter.api.assertThrows

class AssertionTest : FreeSpec({
    "verify response" - {
        "status" - {
            "passes" {
                val response = Response(
                    200,
                    null,
                    null
                )

                verifyResponse(response, 200, null, null)
            }

            "fails" {
                val response = Response(
                    200,
                    null,
                    null
                )

                assertThrows<AssertionError> { verifyResponse(response, 201, null, null) }
            }
        }

        "with body" - {
            "passes" {
                val response = Response(
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
                val response = Response(
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
            "passes" {
                val response = Response(
                    200,
                    mapOf(
                        "key1" to "value1",
                        "key2" to "value2"
                    ),
                    null
                )

                verifyResponse(response, 200, null, mapOf(
                    "key1" to "value1",
                    "key2" to "value2"
                ))
            }

            "fails" {
                val response = Response(
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
        }
    }
})
