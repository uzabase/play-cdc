package com.uzabase.playcdc.internal

import com.uzabase.playcdc.internal.legacy.Request
import com.uzabase.playcdc.internal.legacy.Response

interface Writer {
    fun setup()
    fun write(request: Request)
    fun write(response: Response)
}
