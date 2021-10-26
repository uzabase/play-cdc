package uzabase.com

import io.ktor.application.*
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.http.*
import io.ktor.request.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import uzabase.com.plugins.*


fun main() {
    embeddedServer(Netty, port = 18080, host = "0.0.0.0") {
        configureRouting()

        val client = HttpClient( CIO )

        intercept(ApplicationCallPipeline.Call) {
            println(call.request.uri)

            val proxiedHeaders = call.request.headers
            println(proxiedHeaders.entries().map { Pair(it.key, it.value.joinToString(",")) })

            val proxiedBody = call.receiveText()
            println(proxiedBody)

            call.request.queryParameters.entries().forEach { println("${it.key} ${it.value}") }

            client.request<String> {
                url("https://www.uzabase.com/${call.request.uri}")
                method = call.request.httpMethod
                val headers = call.request.headers.entries().map { Pair(it.key, it.value) }
                headers.forEach { headersOf(it) }
                body = proxiedBody
            }
        }

    }.start(wait = true)
}
