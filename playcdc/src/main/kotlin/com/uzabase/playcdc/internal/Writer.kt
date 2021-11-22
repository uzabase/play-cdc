package com.uzabase.playcdc.internal

interface Writer {
    fun setup()
    fun write(requestJson: RequestJson)
    fun write(responseJson: ResponseJson)
}