package com.uzabase.playcdc.internal.infra

import com.uzabase.playcdc.internal.Response
import org.amshove.kluent.shouldBeEqualTo

fun verifyResponse(response: Response, status: Int, body: String?) {
    response.status shouldBeEqualTo status
    if (body == null) {
        response.jsonBody shouldBeEqualTo emptyMap()
    } else {
        response.jsonBody shouldBeEqualTo toMap(body)
    }
}
