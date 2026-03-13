package ai.learning.sdd;

public final class Plan {
    private final String feature;
    private final String content;

    public Plan(String feature, String content) {
        this.feature = feature;
        this.content = content;
    }

    public String getFeature() { return feature; }
    public String getContent() { return content; }
}
