package com.uzabase.playcdc.internal.infra

import com.fasterxml.jackson.databind.DeserializationFeature
import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper

private val MAPPER = jacksonObjectMapper().configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false)
private val PRETTY_PRINTING_MAPPER = ObjectMapper().writerWithDefaultPrettyPrinter()

@Suppress("UNCHECKED_CAST")
fun toMap(jsonString: String) = MAPPER.readValue(jsonString, Map::class.java) as Map<String, Any>

fun <T> toObject(jsonString: String, clazz: Class<T>): T = MAPPER.readValue(jsonString, clazz)

fun Map<String,Any>.toJsonString(): String = MAPPER.writeValueAsString(this)

fun toPrettyString(obj: Any): String = PRETTY_PRINTING_MAPPER.writeValueAsString(obj)
