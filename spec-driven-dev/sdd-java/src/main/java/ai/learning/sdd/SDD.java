package ai.learning.sdd;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

/**
 * Path conventions and loaders for SDD layout. See docs/sdd/LIBRARY.md.
 */
public final class SDD {

    public static final String SPECS_DIR = "specs";
    public static final String PLANS_DIR = "plans";
    public static final String PROGRESS_DIR = "progress";

    private SDD() {}

    public static Path specPath(String baseDir, String feature) {
        return Paths.get(baseDir, SPECS_DIR, feature + ".md");
    }

    public static Path planPath(String baseDir, String feature) {
        return Paths.get(baseDir, PLANS_DIR, feature + "-plan.md");
    }

    public static Path tasksPath(String baseDir, String feature) {
        return Paths.get(baseDir, PLANS_DIR, feature + "-tasks.md");
    }

    public static Path progressPath(String baseDir, String feature) {
        return Paths.get(baseDir, PROGRESS_DIR, feature + "-progress.md");
    }

    public static PathsResult paths(String baseDir) {
        return new PathsResult(
                Paths.get(baseDir, SPECS_DIR),
                Paths.get(baseDir, PLANS_DIR),
                Paths.get(baseDir, PROGRESS_DIR)
        );
    }

    public static Spec loadSpec(String baseDir, String feature) throws IOException {
        Path p = specPath(baseDir, feature);
        String content = Files.readString(p, StandardCharsets.UTF_8);
        return new Spec(feature, content);
    }

    public static Plan loadPlan(String baseDir, String feature) throws IOException {
        Path p = planPath(baseDir, feature);
        String content = Files.readString(p, StandardCharsets.UTF_8);
        return new Plan(feature, content);
    }

    public static String loadTasks(String baseDir, String feature) throws IOException {
        return Files.readString(tasksPath(baseDir, feature), StandardCharsets.UTF_8);
    }

    public static Progress loadProgress(String baseDir, String feature) throws IOException {
        Path p = progressPath(baseDir, feature);
        String content = Files.readString(p, StandardCharsets.UTF_8);
        return new Progress(feature, content);
    }

    public static final class PathsResult {
        private final Path specsDir;
        private final Path plansDir;
        private final Path progressDir;

        public PathsResult(Path specsDir, Path plansDir, Path progressDir) {
            this.specsDir = specsDir;
            this.plansDir = plansDir;
            this.progressDir = progressDir;
        }

        public Path getSpecsDir() { return specsDir; }
        public Path getPlansDir() { return plansDir; }
        public Path getProgressDir() { return progressDir; }
    }
}
