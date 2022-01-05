package com.uzabase.playcdc.internal.infra

import com.uzabase.playcdc.internal.legacy.Request
import io.kotest.core.spec.style.FreeSpec
import io.kotest.matchers.shouldBe

class RequestTest: FreeSpec({

    "Retrieve header" {
        "Get content-type" {
            val req = Request(
                url = "http://localhost",
                method = "GET",
                headers = mapOf("content-type" to "text-plain"),
                body = emptyMap()
            )

            req.contentType shouldBe "text-plain"
        }

        "Empty content-type" {
            val req = Request(
                url = "http://localhost",
                method = "GET",
                headers = emptyMap(),
                body = emptyMap()
            )

            req.contentType shouldBe null
        }
    }
})