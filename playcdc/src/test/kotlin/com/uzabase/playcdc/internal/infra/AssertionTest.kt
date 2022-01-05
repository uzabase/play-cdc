package com.uzabase.playcdc.internal.infra

import com.uzabase.playcdc.internal.legacy.Response
import io.kotest.core.spec.style.FreeSpec

class AssertionTest : FreeSpec({
    "verify response" - {
        "with body" {
            val response = Response(
                200,
                emptyMap(),
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

            verifyResponse(response, 200, body)
        }

        "without body" {
            val response = Response(
                200,
                emptyMap(),
                emptyMap()
            )

            verifyResponse(response, 200, null)
        }
    }
})
