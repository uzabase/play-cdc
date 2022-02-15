package com.uzabase.playcdc.internal

import com.uzabase.playcdc.internal.Contract.Request
import io.kotest.assertions.throwables.shouldThrow
import io.kotest.core.spec.style.StringSpec
import io.kotest.matchers.should
import io.kotest.matchers.string.startWith

class ContractTest : StringSpec({
  "Request requires either url or urlPath" {
    val exception = shouldThrow<RuntimeException> {
      Request(
        url = null,
        urlPath = null,
        queryParameters = null,
        method = "GET",
        headers = null
      )
    }
    exception.message should startWith("Either `url` or `urlPath` is required")
  }
})
