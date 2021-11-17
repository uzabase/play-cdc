package com.uzabase.playcdc.internal

import com.fasterxml.jackson.databind.ObjectMapper
import java.nio.file.Files
import java.nio.file.Path

class FileWriter(private val folderPath: Path) : Writer {
    private val mapper = ObjectMapper().writerWithDefaultPrettyPrinter()

    override fun createDirectory() {
        Files.createDirectories(folderPath)
    }

    override fun write(requestJson: RequestJson) {
        val jsonString = mapper.writeValueAsString(requestJson)

        Files.write(folderPath.resolve("request.json"), jsonString.toByteArray())
    }

    override fun write(responseJson: ResponseJson) {
        val jsonString = mapper.writeValueAsString(responseJson)

        Files.write(folderPath.resolve("response.json"), jsonString.toByteArray())
    }
}
