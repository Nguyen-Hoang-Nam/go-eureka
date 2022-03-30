package spring.spring.request;

public class SpringRequest {

    private String name;

    public SpringRequest() {}

    /**
     * @param name
     */
    public SpringRequest(String name) {
        this.name = name;
    }

    /**
     * @return the name
     */
    public String getName() {
        return name;
    }

    /**
     * @param name the name to set
     */
    public void setName(String name) {
        this.name = name;
    }
}
