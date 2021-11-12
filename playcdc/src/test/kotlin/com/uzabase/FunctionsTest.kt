package com.uzabase

import com.github.tomakehurst.wiremock.client.WireMock
import io.kotest.core.spec.style.FreeSpec
import io.kotest.matchers.shouldBe
import io.mockk.clearAllMocks

class FunctionsTest : FreeSpec({
    afterTest {
        clearAllMocks()
    }

    "リクエストパスを記録する" {
        var result: String? = null
        val mappingBuilder = WireMock.get("/test")

        storeMock(mappingBuilder, object : Writer {
            override fun writeRequestPath(requestPath: String) {
                result = requestPath
            }
        })

        result shouldBe "/test"
    }
})
