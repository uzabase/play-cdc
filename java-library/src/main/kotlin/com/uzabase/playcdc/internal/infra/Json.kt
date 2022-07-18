package com.uzabase.playcdc.internal.infra

import com.fasterxml.jackson.databind.DeserializationFeature
import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper

private val MAPPER = jacksonObjectMapper().configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false)

@Suppress("UNCHECKED_CAST")
fun toMap(jsonString: String) = MAPPER.readValue(jsonString, Map::class.java) as Map<String, Any>

fun <T> toObject(jsonString: String, clazz: Class<T>): T = MAPPER.readValue(jsonString, clazz)
