import com.github.tomakehurst.wiremock.client.MappingBuilder
import com.thoughtworks.gauge.Step;
import com.github.tomakehurst.wiremock.client.WireMock
import com.thoughtworks.gauge.BeforeScenario

class SampleStepImplementation {
    private val companyApi = WireMock("localhost", 8080)

    @Step("sample step")
    fun sampleStep() {
        println("sample!!")
    }

    @BeforeScenario(tags = ["dependsOnCompanyApi"])
    fun setupCompanyApi() {
        val mappingBuilder = WireMock.get("/test")
            .willReturn(
                WireMock.aResponse()
                    .withHeader("content-type", "text/plain")
                    .withBody("hello")
            )

        companyApi.register(mappingBuilder)
        storeMock(mappingBuilder)
    }

    private fun storeMock(mappingBuilder: MappingBuilder) {
        val methodName = StackWalker.getInstance().walk { frames -> frames.skip(1).findFirst().map { it.methodName } }

        val method = SampleStepImplementation::class.java.getMethod(methodName.get());
        val beforeScenario = method.getAnnotation(BeforeScenario::class.java)
        val tag = beforeScenario.tags[0]

        val req = mappingBuilder.build().request
        val res = mappingBuilder.build().response
        println("$tag: url: ${req.url}, header: ${res.headers}, body: ${res.body}")
    }
}
