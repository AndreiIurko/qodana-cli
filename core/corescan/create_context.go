/*
 * Copyright 2021-2024 JetBrains s.r.o.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package corescan

import (
	"github.com/JetBrains/qodana-cli/v2024/core/startup"
	"github.com/JetBrains/qodana-cli/v2024/platform/cmd"
	"github.com/JetBrains/qodana-cli/v2024/platform/platforminit"
	"github.com/JetBrains/qodana-cli/v2024/platform/qdenv"
	"github.com/JetBrains/qodana-cli/v2024/platform/qdyaml"
	"path/filepath"
	"strings"
)

func CreateContext(
	cliOptions platformcmd.CliOptions,
	initArgs platforminit.Args,
	preparedHost startup.PreparedHost,
	qodanaYaml qdyaml.QodanaYaml,
) Context {
	coverageDir := cliOptions.CoverageDir
	if coverageDir == "" {
		if qdenv.IsContainer() {
			coverageDir = "/data/coverage"
		} else {
			coverageDir = filepath.Join(initArgs.ProjectDir, ".qodana", "code-coverage")
		}
	}

	commit := cliOptions.Commit
	if strings.HasPrefix(commit, "CI") {
		commit = strings.TrimPrefix(commit, "CI")
	}

	return ContextBuilder{
		Linter:                    initArgs.Linter,
		Ide:                       initArgs.Ide,
		Id:                        initArgs.Id,
		IdeDir:                    preparedHost.IdeDir,
		QodanaYaml:                qodanaYaml,
		Prod:                      preparedHost.Prod,
		QodanaToken:               preparedHost.QodanaToken,
		QodanaLicenseOnlyToken:    initArgs.QodanaLicenseOnlyToken,
		ProjectDir:                initArgs.ProjectDir,
		ResultsDir:                initArgs.ResultsDir,
		ConfigDir:                 initArgs.ConfDirPath(),
		LogDir:                    initArgs.LogDir(),
		QodanaSystemDir:           initArgs.QodanaSystemDir,
		CacheDir:                  initArgs.CacheDir,
		ReportDir:                 initArgs.ReportDir,
		CoverageDir:               coverageDir,
		SourceDirectory:           cliOptions.SourceDirectory,
		Env:                       cliOptions.Env_,
		DisableSanity:             cliOptions.DisableSanity,
		ProfileName:               cliOptions.ProfileName,
		ProfilePath:               cliOptions.ProfilePath,
		RunPromo:                  cliOptions.RunPromo,
		StubProfile:               cliOptions.StubProfile,
		Baseline:                  cliOptions.Baseline,
		BaselineIncludeAbsent:     cliOptions.BaselineIncludeAbsent,
		SaveReport:                cliOptions.SaveReport,
		ShowReport:                cliOptions.ShowReport,
		Port:                      cliOptions.Port,
		Property:                  cliOptions.Property,
		Script:                    cliOptions.Script,
		FailThreshold:             cliOptions.FailThreshold,
		Commit:                    commit,
		DiffStart:                 cliOptions.DiffStart,
		DiffEnd:                   cliOptions.DiffEnd,
		ForceLocalChangesScript:   cliOptions.ForceLocalChangesScript,
		AnalysisId:                cliOptions.AnalysisId,
		Volumes:                   cliOptions.Volumes,
		User:                      cliOptions.User,
		PrintProblems:             cliOptions.PrintProblems,
		GenerateCodeClimateReport: cliOptions.GenerateCodeClimateReport,
		SendBitBucketInsights:     cliOptions.SendBitBucketInsights,
		SkipPull:                  cliOptions.SkipPull,
		ClearCache:                initArgs.IsClearCache,
		ConfigName:                cliOptions.ConfigName,
		FullHistory:               cliOptions.FullHistory,
		ApplyFixes:                cliOptions.ApplyFixes,
		Cleanup:                   cliOptions.Cleanup,
		FixesStrategy:             cliOptions.FixesStrategy,
		NoStatistics:              cliOptions.NoStatistics,
		CdnetSolution:             cliOptions.CdnetSolution,
		CdnetProject:              cliOptions.CdnetProject,
		CdnetConfiguration:        cliOptions.CdnetConfiguration,
		CdnetPlatform:             cliOptions.CdnetPlatform,
		CdnetNoBuild:              cliOptions.CdnetNoBuild,
		ClangCompileCommands:      cliOptions.ClangCompileCommands,
		ClangArgs:                 cliOptions.ClangArgs,
		AnalysisTimeoutMs:         cliOptions.AnalysisTimeoutMs,
		AnalysisTimeoutExitCode:   cliOptions.AnalysisTimeoutExitCode,
		JvmDebugPort:              cliOptions.JvmDebugPort,
	}.Build()
}
