package com.uzabase.playcdc.internal

import com.thoughtworks.gauge.BeforeScenario

fun getBasePath(): String = System.getenv("PLAY_CDC_BASE_PATH") ?: "/tmp"

fun getFolderName(): String? = StackWalker.getInstance(StackWalker.Option.RETAIN_CLASS_REFERENCE)
    .walk { frames ->
        frames.map { f ->
            val methods = Class.forName(f.className).methods.filterNotNull().filter { it.name == f.methodName }
            methods.mapNotNull { it.getAnnotation(BeforeScenario::class.java) }.map { it.tags.toList() }.flatten()
        }
            .filter { it.isNotEmpty() }
            .map { it.joinToString("_") }
            .findFirst().orElse(null)
    }
