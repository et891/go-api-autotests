package tests

import (
	"time"

	"github.com/Nikita-Filonov/axiom"
)

func durationPlugin() axiom.Plugin {
	return func(cfg *axiom.Config) {
		cfg.Runtime.EmitTestWrap(func(next axiom.TestAction) axiom.TestAction {
			return func(cfg *axiom.Config) {
				startedAt := time.Now()

				defer func() {
					cfg.T().Logf(
						"[Plugin] test %q finished in %s",
						cfg.Case.Name,
						time.Since(startedAt),
					)
				}()

				next(cfg)
			}
		})

		cfg.Runtime.EmitStepWrap(func(name string, next axiom.StepAction) axiom.StepAction {
			return func() {
				startedAt := time.Now()

				defer func() {
					cfg.T().Logf(
						"[Plugin] step %q finished in %s",
						name,
						time.Since(startedAt),
					)
				}()

				next()
			}
		})
	}
}

func stepCounterPlugin() axiom.Plugin {
	return func(cfg *axiom.Config) {
		stepCount := 0

		cfg.Runtime.EmitStepWrap(func(name string, next axiom.StepAction) axiom.StepAction {
			return func() {
				stepCount++

				cfg.T().Logf(
					"[Plugin] step #%d started: %s",
					stepCount,
					name,
				)

				next()
			}
		})

		cfg.Runtime.EmitTestWrap(func(next axiom.TestAction) axiom.TestAction {
			return func(cfg *axiom.Config) {
				next(cfg)

				cfg.T().Logf(
					"[Plugin] test %q executed %d steps",
					cfg.Case.Name,
					stepCount,
				)
			}
		})
	}
}
