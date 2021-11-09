package com.uzabase

import com.thoughtworks.gauge.BeforeScenario

fun storeMock(writer: (String) -> Unit) {

    val instance = StackWalker.getInstance(StackWalker.Option.RETAIN_CLASS_REFERENCE)
    val frame = instance.walk { frames ->
        frames.filter { f ->
            val methods = Class.forName(f.className).methods.filter { it.name == f.methodName }
            methods.any { it!!.getAnnotation(BeforeScenario::class.java)?.tags?.isNotEmpty() ?: false }
        }.findFirst()
    }.get()

    val methods = Class.forName(frame.className).methods.filter { it.name == frame.methodName }
    methods.mapNotNull { it!!.getAnnotation(BeforeScenario::class.java) }
        .forEach {
            if (it.tags.isNotEmpty()) {
                val tag = it.tags[0]
                writer(tag)
            }
        }
}
