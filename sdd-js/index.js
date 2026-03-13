const path = require('path');
const fs = require('fs');

const SPECS_DIR = 'specs';
const PLANS_DIR = 'plans';
const PROGRESS_DIR = 'progress';

function specPath(baseDir, feature) {
  return path.join(baseDir, SPECS_DIR, `${feature}.md`);
}

function planPath(baseDir, feature) {
  return path.join(baseDir, PLANS_DIR, `${feature}-plan.md`);
}

function tasksPath(baseDir, feature) {
  return path.join(baseDir, PLANS_DIR, `${feature}-tasks.md`);
}

function progressPath(baseDir, feature) {
  return path.join(baseDir, PROGRESS_DIR, `${feature}-progress.md`);
}

function paths(baseDir) {
  return {
    specsDir: path.join(baseDir, SPECS_DIR),
    plansDir: path.join(baseDir, PLANS_DIR),
    progressDir: path.join(baseDir, PROGRESS_DIR),
  };
}

function loadSpec(baseDir, feature) {
  const content = fs.readFileSync(specPath(baseDir, feature), 'utf8');
  return { feature, content };
}

function loadPlan(baseDir, feature) {
  const content = fs.readFileSync(planPath(baseDir, feature), 'utf8');
  return { feature, content };
}

function loadTasks(baseDir, feature) {
  return fs.readFileSync(tasksPath(baseDir, feature), 'utf8');
}

function loadProgress(baseDir, feature) {
  const content = fs.readFileSync(progressPath(baseDir, feature), 'utf8');
  return { feature, content };
}

function loadSpecAsync(baseDir, feature) {
  return fs.promises.readFile(specPath(baseDir, feature), 'utf8').then((content) => ({ feature, content }));
}

function loadPlanAsync(baseDir, feature) {
  return fs.promises.readFile(planPath(baseDir, feature), 'utf8').then((content) => ({ feature, content }));
}

function loadTasksAsync(baseDir, feature) {
  return fs.promises.readFile(tasksPath(baseDir, feature), 'utf8');
}

function loadProgressAsync(baseDir, feature) {
  return fs.promises.readFile(progressPath(baseDir, feature), 'utf8').then((content) => ({ feature, content }));
}

module.exports = {
  SPECS_DIR,
  PLANS_DIR,
  PROGRESS_DIR,
  specPath,
  planPath,
  tasksPath,
  progressPath,
  paths,
  loadSpec,
  loadPlan,
  loadTasks,
  loadProgress,
  loadSpecAsync,
  loadPlanAsync,
  loadTasksAsync,
  loadProgressAsync,
};
