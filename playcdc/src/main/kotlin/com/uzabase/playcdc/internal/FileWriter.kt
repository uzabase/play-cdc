package com.uzabase.playcdc.internal

import com.fasterxml.jackson.databind.ObjectMapper
import java.nio.file.Files
import java.nio.file.Path

class FileWriter(private val folderPath: Path) : Writer {
    private val mapper = ObjectMapper().writerWithDefaultPrettyPrinter()

    override fun setup() {
        Files.createDirectories(folderPath)
    }

    override fun write(requestJson: Request) {
        val jsonString = mapper.writeValueAsString(requestJson)

        Files.write(folderPath.resolve("request.json"), jsonString.toByteArray())
    }

    override fun write(response: Response) {
        val jsonString = mapper.writeValueAsString(response)

        Files.write(folderPath.resolve("response.json"), jsonString.toByteArray())
    }
}
