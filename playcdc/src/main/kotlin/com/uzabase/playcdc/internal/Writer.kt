package com.uzabase.playcdc.internal

interface Writer {
    fun setup()
    fun write(requestJson: Request)
    fun write(response: Response)
}