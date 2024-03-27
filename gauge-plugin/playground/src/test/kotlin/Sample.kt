import com.github.tomakehurst.wiremock.client.WireMock
import com.thoughtworks.gauge.BeforeScenario
import com.thoughtworks.gauge.BeforeSuite
import com.thoughtworks.gauge.Step

class Sample {
    private val wireMock = WireMock(8080)

    @BeforeSuite
    fun loadWiremockMappings() {
        wireMock.loadMappingsFrom("fixtures")
    }

    @BeforeScenario
    fun resetRequests() {
        wireMock.resetRequests()
    }

    @Step("dummy step")
    fun dummy() {
        println("dummy step executed")
    }
}
