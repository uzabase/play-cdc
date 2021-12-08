package com.uzabase.playcdc.internal

interface Writer {
    fun setup()
    fun write(request: Request)
    fun write(response: Response)
}
