export const SPECS_DIR: string;
export const PLANS_DIR: string;
export const PROGRESS_DIR: string;

export function specPath(baseDir: string, feature: string): string;
export function planPath(baseDir: string, feature: string): string;
export function tasksPath(baseDir: string, feature: string): string;
export function progressPath(baseDir: string, feature: string): string;

export interface PathsResult {
  specsDir: string;
  plansDir: string;
  progressDir: string;
}
export function paths(baseDir: string): PathsResult;

export interface DocResult {
  feature: string;
  content: string;
}
export function loadSpec(baseDir: string, feature: string): DocResult;
export function loadPlan(baseDir: string, feature: string): DocResult;
export function loadTasks(baseDir: string, feature: string): string;
export function loadProgress(baseDir: string, feature: string): DocResult;

export function loadSpecAsync(baseDir: string, feature: string): Promise<DocResult>;
export function loadPlanAsync(baseDir: string, feature: string): Promise<DocResult>;
export function loadTasksAsync(baseDir: string, feature: string): Promise<string>;
export function loadProgressAsync(baseDir: string, feature: string): Promise<DocResult>;
