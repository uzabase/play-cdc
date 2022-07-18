package com.uzabase.playcdc.internal.infra

import com.uzabase.playcdc.internal.Contract
import io.kotest.core.spec.style.FreeSpec
import io.kotest.matchers.shouldBe

class HttpTest : FreeSpec({
    "queryParameters" - {
        "single" {
            val request = Contract.Request(
                url = null,
                urlPath = "/path",
                queryParameters = mapOf(
                    "key1" to mapOf("equalTo" to "value1")
                ),
                method = "GET",
                headers = null
            )
            queryParameters(request) shouldBe "?key1=value1"
        }

        "multiple" {
            val request = Contract.Request(
                url = null,
                urlPath = "/path",
                queryParameters = mapOf(
                    "key1" to mapOf("equalTo" to "value1"),
                    "key2" to mapOf("equalTo" to "value2")
                ),
                method = "GET",
                headers = null
            )
            queryParameters(request) shouldBe "?key1=value1&key2=value2"
        }

        "none" {
            val request = Contract.Request(
                url = "/path",
                urlPath = null,
                queryParameters = null,
                method = "GET",
                headers = null
            )
            queryParameters(request) shouldBe ""
        }
    }

    "toArray" {
        toArray(mapOf("key" to mapOf("equalTo" to "value"))) shouldBe arrayOf("key", "value")
    }
})
