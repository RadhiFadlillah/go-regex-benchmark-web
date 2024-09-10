package main

import (
	"github.com/dlclark/regexp2"
	"github.com/dlclark/regexp2/helpers"
	"github.com/dlclark/regexp2/syntax"
	"unicode"
)

/*
Capture(index = 0, unindex = -1)
 Concatenate
  SetloopAtomic(Set = [\+-\.\w])(Min = 1, Max = inf)
  UpdateBumpalong
  One(Ch = @)
  Setloop(Set = [-\.\w])(Min = 1, Max = inf)
  One(Ch = \.)
  SetloopAtomic(Set = [-\.\w])(Min = 1, Max = inf)
*/
// From main.go:14:34
// Pattern: "[\\w\\.+-]+@[\\w\\.-]+\\.[\\w\\.-]+"
// Options: regexp2.None
type rxEmail_Engine struct{}

func (rxEmail_Engine) Caps() map[int]int        { return nil }
func (rxEmail_Engine) CapNames() map[string]int { return nil }
func (rxEmail_Engine) CapsList() []string       { return nil }
func (rxEmail_Engine) CapSize() int             { return 1 }

func (rxEmail_Engine) FindFirstChar(r *regexp2.Runner) bool {
	pos := r.Runtextpos
	// Any possible match is at least 5 characters
	if pos <= len(r.Runtext)-5 {
		// The pattern begins with an atomic loop for [\+-\.\w] {DescribeSet(target.LoopNode.Str!)}, followed by the character '@'
		// Search for the literal, and then walk backwards to the beginning of the loop.
		for {
			slice := r.Runtext[pos:]

			i := helpers.IndexOfAny1(slice, '@')
			if i < 0 {
				break
			}

			prev := i - 1
			for uint(prev) < uint(len(slice)) && set_f8012346c36030544407c813dd9113ad6c1c8080037c9de30fcfa701bc6d3567.CharIn(slice[prev]) {
				prev--
			}

			if (i - prev - 1) < 1 {
				pos += i + 1
				continue
			}

			r.Runtextpos = pos + prev + 1
			r.Runtrackpos = pos + i
			return true

		}
	}

	// No match found
	r.Runtextpos = len(r.Runtext)
	return false
}

func (rxEmail_Engine) Execute(r *regexp2.Runner) error {
	var charloop_starting_pos, charloop_ending_pos = 0, 0
	iteration := 0
	iteration1 := 0
	pos := r.Runtextpos
	matchStart := pos

	var slice = r.Runtext[pos:]

	// Node: Concatenate
	// Skip loop already matched in TryFindNextPossibleStartingPosition.
	pos = r.Runtrackpos
	slice = r.Runtext[pos:]

	// Node: UpdateBumpalong
	// Advance the next matching position.
	if r.Runtextpos < pos {
		r.Runtextpos = pos
	}

	// Node: One(Ch = @)
	// Match '@'.
	if len(slice) == 0 || slice[0] != '@' {
		return nil // The input didn't match.
	}

	// Node: Setloop(Set = [-\.\w])(Min = 1, Max = inf)
	// Match [-\.\w] greedily at least once.
	pos++
	slice = r.Runtext[pos:]
	charloop_starting_pos = pos

	iteration = 0
	for iteration < len(slice) && set_bc507dd5628b7f62f93ecd744fb34f2123f824a49c4534ea1e4a3bc3e53aed33.CharIn(slice[iteration]) {
		iteration++
	}

	if iteration == 0 {
		return nil // The input didn't match.
	}

	slice = slice[iteration:]
	pos += iteration

	charloop_ending_pos = pos
	charloop_starting_pos++
	goto CharLoopEnd

CharLoopBacktrack:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos >= charloop_ending_pos {
		return nil // The input didn't match.
	}
	charloop_ending_pos = helpers.LastIndexOfAny1(r.Runtext[charloop_starting_pos:charloop_ending_pos], '.')
	if charloop_ending_pos < 0 { // miss
		return nil // The input didn't match.
	}
	charloop_ending_pos += charloop_starting_pos
	pos = charloop_ending_pos
	slice = r.Runtext[pos:]

CharLoopEnd:

	// Node: One(Ch = \.)
	// Match '.'.
	if len(slice) == 0 || slice[0] != '.' {
		goto CharLoopBacktrack
	}

	// Node: SetloopAtomic(Set = [-\.\w])(Min = 1, Max = inf)
	// Match [-\.\w] atomically at least once.
	pos++
	slice = r.Runtext[pos:]
	iteration1 = 0
	for iteration1 < len(slice) && set_bc507dd5628b7f62f93ecd744fb34f2123f824a49c4534ea1e4a3bc3e53aed33.CharIn(slice[iteration1]) {
		iteration1++
	}

	if iteration1 == 0 {
		goto CharLoopBacktrack
	}

	slice = slice[iteration1:]
	pos += iteration1

	// The input matched.
	r.Runtextpos = pos
	r.Capture(0, matchStart, pos)
	// just to prevent an unused var error in certain regex's
	var _ = slice
	return nil
}

/*
Capture(index = 0, unindex = -1)
 Concatenate
  SetloopAtomic(Set = [\w])(Min = 1, Max = inf)
  UpdateBumpalong
  Multi(String = "://")
  Setloop(Set = [^\#/\?\s])(Min = 1, Max = inf)
  Setloop(Set = [^\#\?\s])(Min = 1, Max = inf)
  Loop(Min = 0, Max = 1)
   Concatenate
    One(Ch = \?)
    Setloop(Set = [^\#\s])(Min = 0, Max = inf)
  Atomic
   Loop(Min = 0, Max = 1)
    Concatenate
     One(Ch = \#)
     SetloopAtomic(Set = [^\s])(Min = 0, Max = inf)
*/
// From main.go:15:34
// Pattern: "[\\w]+://[^/\\s?#]+[^\\s?#]+(?:\\?[^\\s#]*)?(?:#[^\\s]*)?"
// Options: regexp2.None
type rxURI_Engine struct{}

func (rxURI_Engine) Caps() map[int]int        { return nil }
func (rxURI_Engine) CapNames() map[string]int { return nil }
func (rxURI_Engine) CapsList() []string       { return nil }
func (rxURI_Engine) CapSize() int             { return 1 }

func (rxURI_Engine) FindFirstChar(r *regexp2.Runner) bool {
	pos := r.Runtextpos
	// Any possible match is at least 6 characters
	if pos <= len(r.Runtext)-6 {
		// The pattern begins with an atomic loop for [\w] {DescribeSet(target.LoopNode.Str!)}, followed by the string ://
		// Search for the literal, and then walk backwards to the beginning of the loop.
		for {
			slice := r.Runtext[pos:]

			i := helpers.IndexOf(slice, []rune("://"))
			if i < 0 {
				break
			}

			prev := i - 1
			for uint(prev) < uint(len(slice)) && helpers.IsWordChar(slice[prev]) {
				prev--
			}

			if (i - prev - 1) < 1 {
				pos += i + 1
				continue
			}

			r.Runtextpos = pos + prev + 1
			r.Runtrackpos = pos + i
			return true

		}
	}

	// No match found
	r.Runtextpos = len(r.Runtext)
	return false
}

func (rxURI_Engine) Execute(r *regexp2.Runner) error {
	var charloop_starting_pos, charloop_ending_pos = 0, 0
	iteration := 0
	var charloop_starting_pos1, charloop_ending_pos1 = 0, 0
	iteration1 := 0
	loop_iteration := 0
	var charloop_starting_pos2, charloop_ending_pos2 = 0, 0
	iteration2 := 0
	loop_iteration1 := 0
	iteration3 := 0
	startingStackpos := 0
	pos := r.Runtextpos
	matchStart := pos

	var slice = r.Runtext[pos:]

	// Node: Concatenate
	// Skip loop already matched in TryFindNextPossibleStartingPosition.
	pos = r.Runtrackpos
	slice = r.Runtext[pos:]

	// Node: UpdateBumpalong
	// Advance the next matching position.
	if r.Runtextpos < pos {
		r.Runtextpos = pos
	}

	// Node: Multi(String = "://")
	// Match the string "://".
	if !helpers.StartsWith(slice, []rune("://")) {
		return nil // The input didn't match.
	}

	// Node: Setloop(Set = [^\#/\?\s])(Min = 1, Max = inf)
	// Match [^\#/\?\s] greedily at least once.
	pos += 3
	slice = r.Runtext[pos:]
	charloop_starting_pos = pos

	iteration = 0
	for iteration < len(slice) && set_201e66754544c28afd79d8418937f5372d668f4011ef28456fbb0936ea3e0238.CharIn(slice[iteration]) {
		iteration++
	}

	if iteration == 0 {
		return nil // The input didn't match.
	}

	slice = slice[iteration:]
	pos += iteration

	charloop_ending_pos = pos
	charloop_starting_pos++
	goto CharLoopEnd

CharLoopBacktrack:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos >= charloop_ending_pos {
		return nil // The input didn't match.
	}
	charloop_ending_pos--
	pos = charloop_ending_pos
	slice = r.Runtext[pos:]

CharLoopEnd:

	// Node: Setloop(Set = [^\#\?\s])(Min = 1, Max = inf)
	// Match [^\#\?\s] greedily at least once.
	charloop_starting_pos1 = pos

	iteration1 = 0
	for iteration1 < len(slice) && set_d224ef67dc0f70cf5e774cc9c41501d7c305e2a2d786eb4c19dc0484e59b68fd.CharIn(slice[iteration1]) {
		iteration1++
	}

	if iteration1 == 0 {
		goto CharLoopBacktrack
	}

	slice = slice[iteration1:]
	pos += iteration1

	charloop_ending_pos1 = pos
	charloop_starting_pos1++
	goto CharLoopEnd1

CharLoopBacktrack1:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos1 >= charloop_ending_pos1 {
		goto CharLoopBacktrack
	}
	charloop_ending_pos1 = helpers.LastIndexOfAny1(r.Runtext[charloop_starting_pos1:charloop_ending_pos1], '?')
	if charloop_ending_pos1 < 0 { // miss
		goto CharLoopBacktrack
	}
	charloop_ending_pos1 += charloop_starting_pos1
	pos = charloop_ending_pos1
	slice = r.Runtext[pos:]

CharLoopEnd1:

	// Node: Loop(Min = 0, Max = 1)
	// Optional (greedy).
	loop_iteration = 0

LoopBody:
	r.StackPush(pos)

	loop_iteration++

	// Node: Concatenate
	// Node: One(Ch = \?)
	// Match '?'.
	if len(slice) == 0 || slice[0] != '?' {
		goto LoopIterationNoMatch
	}

	// Node: Setloop(Set = [^\#\s])(Min = 0, Max = inf)
	// Match [^\#\s] greedily any number of times.
	pos++
	slice = r.Runtext[pos:]
	charloop_starting_pos2 = pos

	iteration2 = 0
	for iteration2 < len(slice) && set_0811f1ec56483a6ff109ad6cdaabc3ae0d433e896a4b209895d649eca8b8e591.CharIn(slice[iteration2]) {
		iteration2++
	}

	slice = slice[iteration2:]
	pos += iteration2

	charloop_ending_pos2 = pos
	goto CharLoopEnd2

CharLoopBacktrack2:
	charloop_ending_pos2 = r.StackPop()
	charloop_starting_pos2 = r.StackPop()

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos2 >= charloop_ending_pos2 {
		goto LoopIterationNoMatch
	}
	charloop_ending_pos2--
	pos = charloop_ending_pos2
	slice = r.Runtext[pos:]

CharLoopEnd2:
	r.StackPush2(charloop_starting_pos2, charloop_ending_pos2)

	// The loop has an upper bound of 1. Continue iterating greedily if it hasn't yet been reached.
	if loop_iteration == 0 {
		goto LoopBody
	}
	goto LoopEnd

	// The loop iteration failed. Put state back to the way it was before the iteration.
LoopIterationNoMatch:
	loop_iteration--
	if loop_iteration < 0 {
		// Unable to match the remainder of the expression after exhausting the loop.
		goto CharLoopBacktrack1
	}
	pos = r.StackPop()
	slice = r.Runtext[pos:]
	goto LoopEnd

LoopBacktrack:
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if loop_iteration == 0 {
		// No iterations of the loop remain to backtrack into. Fail the loop.
		goto CharLoopBacktrack1
	}
	goto CharLoopBacktrack2
LoopEnd:
	;

	// Node: Atomic
	// Node: Loop(Min = 0, Max = 1)
	// Optional (greedy).
	startingStackpos = r.Runstackpos
	loop_iteration1 = 0

LoopBody1:
	;
	r.StackPush(pos)

	loop_iteration1++

	// Node: Concatenate
	// Node: One(Ch = \#)
	// Match '#'.
	if len(slice) == 0 || slice[0] != '#' {
		goto LoopIterationNoMatch1
	}

	// Node: SetloopAtomic(Set = [^\s])(Min = 0, Max = inf)
	// Match [^\s] atomically any number of times.
	iteration3 = 1
	for iteration3 < len(slice) && !unicode.IsSpace(slice[iteration3]) {
		iteration3++
	}

	slice = slice[iteration3:]
	pos += iteration3

	// The loop has an upper bound of 1. Continue iterating greedily if it hasn't yet been reached.
	if loop_iteration1 == 0 {
		goto LoopBody1
	}
	goto LoopEnd1

	// The loop iteration failed. Put state back to the way it was before the iteration.
LoopIterationNoMatch1:
	loop_iteration1--
	if loop_iteration1 < 0 {
		// Unable to match the remainder of the expression after exhausting the loop.
		goto LoopBacktrack
	}
	pos = r.StackPop()
	slice = r.Runtext[pos:]
LoopEnd1:
	r.Runstackpos = startingStackpos // Ensure any remaining backtracking state is removed.

	// The input matched.
	r.Runtextpos = pos
	r.Capture(0, matchStart, pos)
	// just to prevent an unused var error in certain regex's
	var _ = slice
	return nil
}

/*
Capture(index = 0, unindex = -1)
 Concatenate
  Loop(Min = 3, Max = 3)
   Concatenate
    Alternate
     Concatenate
      One(Ch = 2)
      Alternate
       Concatenate
        One(Ch = 5)
        Set(Set = [0-9])
       Concatenate
        Set(Set = [0-4])
        Set(Set = [0-9])
     Concatenate
      Setloop(Set = [01])(Min = 0, Max = 1)
      SetloopAtomic(Set = [0-9])(Min = 2, Max = 2)
    One(Ch = \.)
  Atomic
   Alternate
    Concatenate
     One(Ch = 2)
     Atomic
      Alternate
       Concatenate
        One(Ch = 5)
        Set(Set = [0-5])
       Concatenate
        Set(Set = [0-4])
        Set(Set = [0-9])
    Concatenate
     Setloop(Set = [01])(Min = 0, Max = 1)
     SetloopAtomic(Set = [0-9])(Min = 2, Max = 2)
*/
// From main.go:16:34
// Pattern: "(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])"
// Options: regexp2.None
type rxIP_Engine struct{}

func (rxIP_Engine) Caps() map[int]int        { return nil }
func (rxIP_Engine) CapNames() map[string]int { return nil }
func (rxIP_Engine) CapsList() []string       { return nil }
func (rxIP_Engine) CapSize() int             { return 1 }

func (rxIP_Engine) FindFirstChar(r *regexp2.Runner) bool {
	pos := r.Runtextpos
	// Any possible match is at least 11 characters
	if pos <= len(r.Runtext)-11 {
		// The pattern begins with [0-9]
		// Find the next occurrence. If it can't be found, there's no match.
		i := helpers.IndexOfAnyInRange(r.Runtext[pos:], '0', '9')
		if i >= 0 {
			r.Runtextpos = pos + i
			return true
		}

	}

	// No match found
	r.Runtextpos = len(r.Runtext)
	return false
}

func (rxIP_Engine) Execute(r *regexp2.Runner) error {
	loop_iteration := 0
	alternation_starting_pos := 0
	var charloop_starting_pos, charloop_ending_pos = 0, 0
	atomic_stackpos := 0
	alternation_starting_pos1 := 0
	var charloop_starting_pos1, charloop_ending_pos1 = 0, 0
	pos := r.Runtextpos
	matchStart := pos

	var slice = r.Runtext[pos:]

	// Node: Concatenate
	// Node: Loop(Min = 3, Max = 3)
	// Loop exactly 3 times.
	loop_iteration = 0

LoopBody:
	r.StackPush(pos)

	loop_iteration++

	// Node: Concatenate
	// Node: Alternate
	// Match with 2 alternative expressions.
	alternation_starting_pos = pos

	// Branch 0
	// Node: Concatenate
	// Node: One(Ch = 2)
	// Match '2'.
	if len(slice) == 0 || slice[0] != '2' {
		goto AlternationBranch
	}

	// Node: Alternate
	// Match with 2 alternative expressions.
	if len(slice) < 2 {
		goto AlternationBranch
	}

	switch slice[1] {
	case '5':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [0-9])
		// Match [0-9].
		if len(slice) < 3 || !helpers.IsBetween(slice[2], '0', '9') {
			goto AlternationBranch
		}

		pos += 3
		slice = r.Runtext[pos:]

	case '0', '1', '2', '3', '4':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [0-9])
		// Match [0-9].
		if len(slice) < 3 || !helpers.IsBetween(slice[2], '0', '9') {
			goto AlternationBranch
		}

		pos += 3
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch
	}

	r.StackPush2(0, alternation_starting_pos)
	goto AlternationMatch

AlternationBranch:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]

	// Branch 1
	// Node: Concatenate
	// Node: Setloop(Set = [01])(Min = 0, Max = 1)
	// Match [01] greedily, optionally.
	charloop_starting_pos = pos

	if len(slice) > 0 && helpers.IsBetween(slice[0], '0', '1') {
		slice = slice[1:]
		pos++
	}

	charloop_ending_pos = pos
	goto CharLoopEnd

CharLoopBacktrack:
	charloop_ending_pos = r.StackPop()
	charloop_starting_pos = r.StackPop()

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos >= charloop_ending_pos {
		goto LoopIterationNoMatch
	}
	charloop_ending_pos--
	pos = charloop_ending_pos
	slice = r.Runtext[pos:]

CharLoopEnd:
	r.StackPush2(charloop_starting_pos, charloop_ending_pos)

	// Node: SetloopAtomic(Set = [0-9])(Min = 2, Max = 2)
	// Match [0-9] exactly 2 times.
	if len(slice) < 2 ||
		!helpers.IsBetween(slice[0], '0', '9') ||
		!helpers.IsBetween(slice[1], '0', '9') {
		goto CharLoopBacktrack
	}

	r.StackPush2(1, alternation_starting_pos)
	pos += 2
	slice = r.Runtext[pos:]
	goto AlternationMatch

AlternationBacktrack:
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	alternation_starting_pos = r.StackPop()
	switch r.StackPop() {
	case 0:
		goto AlternationBranch
	case 1:
		goto CharLoopBacktrack
	}

AlternationMatch:
	;

	// Node: One(Ch = \.)
	// Match '.'.
	if len(slice) == 0 || slice[0] != '.' {
		goto AlternationBacktrack
	}

	pos++
	slice = r.Runtext[pos:]

	// The loop has an upper bound of 3. Continue iterating greedily if it hasn't yet been reached.
	if loop_iteration < 3 {
		goto LoopBody
	}
	goto LoopEnd

	// The loop iteration failed. Put state back to the way it was before the iteration.
LoopIterationNoMatch:
	loop_iteration--
	if loop_iteration < 0 {
		// Unable to match the remainder of the expression after exhausting the loop.
		return nil // The input didn't match.
	}
	pos = r.StackPop()
	slice = r.Runtext[pos:]
	if loop_iteration == 0 {
		// No iterations have been matched to backtrack into. Fail the loop.
		return nil // The input didn't match.
	}

	if loop_iteration < 3 {
		// All possible iterations have matched, but it's below the required minimum of 3.
		// Backtrack into the prior iteration.
		goto AlternationBacktrack
	}

	goto LoopEnd

LoopBacktrack:
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if loop_iteration == 0 {
		// No iterations of the loop remain to backtrack into. Fail the loop.
		return nil // The input didn't match.
	}
	goto AlternationBacktrack
LoopEnd:
	;

	// Node: Atomic
	// Atomic group.
	atomic_stackpos = r.Runstackpos

	// Node: Alternate
	// Match with 2 alternative expressions, atomically.
	alternation_starting_pos1 = pos

	// Branch 0
	// Node: Concatenate
	// Node: One(Ch = 2)
	// Match '2'.
	if len(slice) == 0 || slice[0] != '2' {
		goto AlternationBranch1
	}

	// Node: Atomic
	// Node: Alternate
	// Match with 2 alternative expressions, atomically.
	if len(slice) < 2 {
		goto AlternationBranch1
	}

	switch slice[1] {
	case '5':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [0-5])
		// Match [0-5].
		if len(slice) < 3 || !helpers.IsBetween(slice[2], '0', '5') {
			goto AlternationBranch1
		}

		pos += 3
		slice = r.Runtext[pos:]

	case '0', '1', '2', '3', '4':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [0-9])
		// Match [0-9].
		if len(slice) < 3 || !helpers.IsBetween(slice[2], '0', '9') {
			goto AlternationBranch1
		}

		pos += 3
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch1
	}

	goto AlternationMatch1

AlternationBranch1:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]

	// Branch 1
	// Node: Concatenate
	// Node: Setloop(Set = [01])(Min = 0, Max = 1)
	// Match [01] greedily, optionally.
	charloop_starting_pos1 = pos

	if len(slice) > 0 && helpers.IsBetween(slice[0], '0', '1') {
		slice = slice[1:]
		pos++
	}

	charloop_ending_pos1 = pos
	goto CharLoopEnd1

CharLoopBacktrack1:

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos1 >= charloop_ending_pos1 {
		goto LoopBacktrack
	}
	charloop_ending_pos1--
	pos = charloop_ending_pos1
	slice = r.Runtext[pos:]

CharLoopEnd1:

	// Node: SetloopAtomic(Set = [0-9])(Min = 2, Max = 2)
	// Match [0-9] exactly 2 times.
	if len(slice) < 2 ||
		!helpers.IsBetween(slice[0], '0', '9') ||
		!helpers.IsBetween(slice[1], '0', '9') {
		goto CharLoopBacktrack1
	}

	pos += 2
	slice = r.Runtext[pos:]

AlternationMatch1:
	;

	r.Runstackpos = atomic_stackpos

	// The input matched.
	r.Runtextpos = pos
	r.Capture(0, matchStart, pos)
	// just to prevent an unused var error in certain regex's
	var _ = slice
	return nil
}

/*
Capture(index = 0, unindex = -1)
 Atomic
  Alternate
   Concatenate
    Capture(index = 1, unindex = -1)
     Alternate
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Aa])
       Set(Set = [Nn])
       Set(Set = [Uu])
       Set(Set = [Aa])
       Set(Set = [Rr])
       SetloopAtomic(Set = [Yy])(Min = 0, Max = 1)
      Concatenate
       Set(Set = [Ff])
       Set(Set = [Ee])
       Set(Set = [Bb])
       Set(Set = [Rr])
       Set(Set = [Uu])
       Set(Set = [Aa])
       Set(Set = [Rr])
       SetloopAtomic(Set = [Yy])(Min = 0, Max = 1)
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Cc])
       Set(Set = [Hh])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [PVpv])
       Set(Set = [Rr])
       Set(Set = [Ii])
       Set(Set = [Ll])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Aa])
       Set(Set = [IYiy])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Uu])
       Alternate
        Concatenate
         Set(Set = [Nn])
         Set(Set = [EIei])
        Concatenate
         Set(Set = [Ll])
         Set(Set = [IYiy])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Uu])
       Set(Set = [Gg])
       Set(Set = [Uu])
       Set(Set = [Ssſ])
       Set(Set = [Tt])
      Concatenate
       Set(Set = [Ssſ])
       Set(Set = [Ee])
       Set(Set = [Pp])
       Set(Set = [Tt])
       Set(Set = [Ee])
       Set(Set = [Mm])
       Set(Set = [Bb])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Oo])
       Set(Set = [CKckK])
       Set(Set = [Tt])
       Set(Set = [Oo])
       Set(Set = [Bb])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Nn])
       Set(Set = [Oo])
       Set(Set = [Vv])
       Set(Set = [Ee])
       Set(Set = [Mm])
       Set(Set = [Bb])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Dd])
       Set(Set = [Ee])
       Set(Set = [CSZcszſ])
       Set(Set = [Ee])
       Set(Set = [Mm])
       Set(Set = [Bb])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Aa])
       Set(Set = [Nn])
      Concatenate
       Set(Set = [Ff])
       Set(Set = [Ee])
       Set(Set = [Bb])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [AaÄä])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Pp])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Uu])
       Set(Set = [LNln])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Uu])
       Set(Set = [Gg])
      Concatenate
       Set(Set = [Ssſ])
       Set(Set = [Ee])
       Set(Set = [Pp])
      Concatenate
       Set(Set = [Oo])
       Set(Set = [CKckK])
       Set(Set = [Tt])
      Concatenate
       Set(Set = [Nn])
       Set(Set = [Oo])
       Set(Set = [Vv])
      Concatenate
       Set(Set = [Dd])
       Set(Set = [Ee])
       Set(Set = [CZcz])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Aa])
       Set(Set = [Nn])
       Set(Set = [Uu])
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Ii])
      Concatenate
       Set(Set = [Ff])
       Set(Set = [Ee])
       Set(Set = [Bb])
       Set(Set = [Rr])
       Set(Set = [Uu])
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Ii])
      Concatenate
       Set(Set = [Mm])
       Alternate
        Concatenate
         Set(Set = [Aa])
         Set(Set = [Rr])
         Set(Set = [Ee])
         Set(Set = [Tt])
        Concatenate
         Set(Set = [Ee])
         Set(Set = [Ii])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Gg])
       Set(Set = [Uu])
       Set(Set = [Ssſ])
       Set(Set = [Tt])
       Set(Set = [Uu])
       Set(Set = [Ssſ])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Ää])
       SetloopAtomic(Set = [Nn])(Min = 2, Max = 2)
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Ff])
       Set(Set = [Ee])
       Set(Set = [Bb])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Ää])
       Set(Set = [Rr])
       Set(Set = [Zz])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Aa])
       Set(Set = [Nn])
       Set(Set = [Vv])
       Set(Set = [Ii])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Ff])
       Set(Set = [Éé])
       Set(Set = [Vv])
       Set(Set = [Rr])
       Set(Set = [Ii])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Ssſ])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Uu])
       Set(Set = [Ii])
       Alternate
        Set(Set = [Nn])
        Concatenate
         SetloopAtomic(Set = [Ll])(Min = 2, Max = 2)
         Set(Set = [Ee])
         Set(Set = [Tt])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Oo])
       Set(Set = [Uu])
       Set(Set = [Tt])
      Concatenate
       Set(Set = [Ssſ])
       Set(Set = [Ee])
       Set(Set = [Pp])
       Set(Set = [Tt])
       Set(Set = [Ee])
       Set(Set = [Mm])
       Set(Set = [Bb])
       Set(Set = [Rr])
       Set(Set = [Ee])
      Concatenate
       Set(Set = [Oo])
       Set(Set = [Cc])
       Set(Set = [Tt])
       Set(Set = [Oo])
       Set(Set = [Bb])
       Set(Set = [Rr])
       Set(Set = [Ee])
      Concatenate
       Set(Set = [Nn])
       Set(Set = [Oo])
       Set(Set = [Vv])
       Set(Set = [Ee])
       Set(Set = [Mm])
       Set(Set = [Bb])
       Set(Set = [Rr])
       Set(Set = [Ee])
      Concatenate
       Set(Set = [Dd])
       Set(Set = [Éé])
       Set(Set = [Cc])
       Set(Set = [Ee])
       Set(Set = [Mm])
       Set(Set = [Bb])
       Set(Set = [Rr])
       Set(Set = [Ee])
      Concatenate
       Set(Set = [Oo])
       Set(Set = [Cc])
       Set(Set = [Aa])
       Set(Set = [KkK])
      Concatenate
       Set(Set = [Şş])
       Set(Set = [Uu])
       Set(Set = [Bb])
       Set(Set = [Aa])
       Set(Set = [Tt])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Tt])
      Concatenate
       Set(Set = [Nn])
       Set(Set = [Ii])
       Set(Set = [Ssſ])
       Set(Set = [Aa])
       Set(Set = [Nn])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Aa])
       Set(Set = [Yy])
       One(Ch = ı)
       Set(Set = [Ssſ])
      Concatenate
       Set(Set = [Hh])
       Set(Set = [Aa])
       Set(Set = [Zz])
       Set(Set = [Ii])
       Set(Set = [Rr])
       Set(Set = [Aa])
       Set(Set = [Nn])
      Concatenate
       Set(Set = [Tt])
       Set(Set = [Ee])
       SetloopAtomic(Set = [Mm])(Min = 2, Max = 2)
       Set(Set = [Uu])
       Set(Set = [Zz])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Ğğ])
       Set(Set = [Uu])
       Set(Set = [Ssſ])
       Set(Set = [Tt])
       Set(Set = [Oo])
       Set(Set = [Ssſ])
      Concatenate
       Set(Set = [Ee])
       Alternate
        Concatenate
         Set(Set = [Yy])
         Set(Set = [Ll])
         Set(Set = [Üü])
         Set(Set = [Ll])
        Concatenate
         Set(Set = [KkK])
         Set(Set = [Ii])
         Set(Set = [Mm])
      Concatenate
       Set(Set = [KkK])
       Set(Set = [Aa])
       Set(Set = [Ssſ])
       One(Ch = ı)
       Set(Set = [Mm])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Aa])
       Set(Set = [Ll])
       One(Ch = ı)
       Set(Set = [KkK])
      Concatenate
       Set(Set = [Oo])
       Set(Set = [Cc])
       Set(Set = [Aa])
      Concatenate
       Set(Set = [Şş])
       Set(Set = [Uu])
       Set(Set = [Bb])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Aa])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Nn])
       Set(Set = [Ii])
       Set(Set = [Ssſ])
      Concatenate
       Set(Set = [Hh])
       Set(Set = [Aa])
       Set(Set = [Zz])
      Concatenate
       Set(Set = [Tt])
       Set(Set = [Ee])
       Set(Set = [Mm])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Ğğ])
       Set(Set = [Uu])
      Concatenate
       Set(Set = [Ee])
       Alternate
        Concatenate
         Set(Set = [Yy])
         Set(Set = [Ll])
        Concatenate
         Set(Set = [KkK])
         Set(Set = [Ii])
      Concatenate
       Set(Set = [KkK])
       Set(Set = [Aa])
       Set(Set = [Ssſ])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Aa])
    Set(Set = [\s])
    Capture(index = 2, unindex = -1)
     Concatenate
      Setloop(Set = [0-3])(Min = 0, Max = 1)
      Set(Set = [0-9])
    Loop(Min = 0, Max = 1)
     Alternate
      Concatenate
       Set(Set = [Ssſ])
       Set(Set = [Tt])
      Concatenate
       Set(Set = [Nn])
       Set(Set = [Dd])
      Concatenate
       Set(Set = [Rr])
       Set(Set = [Dd])
      Concatenate
       Set(Set = [Tt])
       Set(Set = [Hh])
    OneloopAtomic(Ch = ,)(Min = 0, Max = 1)
    Set(Set = [\s])
    Capture(index = 3, unindex = -1)
     Atomic
      Alternate
       Concatenate
        Multi(String = "199")
        Set(Set = [0-9])
       Concatenate
        Multi(String = "20")
        Set(Set = [0-3])
        Set(Set = [0-9])
   Concatenate
    Capture(index = 4, unindex = -1)
     Concatenate
      Setloop(Set = [0-3])(Min = 0, Max = 1)
      Set(Set = [0-9])
    Loop(Min = 0, Max = 1)
     Alternate
      Concatenate
       Set(Set = [Ssſ])
       Set(Set = [Tt])
      Concatenate
       Set(Set = [Nn])
       Set(Set = [Dd])
      Concatenate
       Set(Set = [Rr])
       Set(Set = [Dd])
      Concatenate
       Set(Set = [Tt])
       Set(Set = [Hh])
      One(Ch = \.)
    Set(Set = [\s])
    Loop(Min = 0, Max = 1)
     Concatenate
      Set(Set = [Oo])
      Set(Set = [Ff])
      Set(Set = [\s])
    Capture(index = 5, unindex = -1)
     Alternate
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Aa])
       Set(Set = [Nn])
       Set(Set = [Uu])
       Set(Set = [Aa])
       Set(Set = [Rr])
       Setloop(Set = [Yy])(Min = 0, Max = 1)
      Concatenate
       Set(Set = [Ff])
       Set(Set = [Ee])
       Set(Set = [Bb])
       Set(Set = [Rr])
       Set(Set = [Uu])
       Set(Set = [Aa])
       Set(Set = [Rr])
       Setloop(Set = [Yy])(Min = 0, Max = 1)
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Cc])
       Set(Set = [Hh])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [PVpv])
       Set(Set = [Rr])
       Set(Set = [Ii])
       Set(Set = [Ll])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Aa])
       Set(Set = [IYiy])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Uu])
       Alternate
        Concatenate
         Set(Set = [Nn])
         Set(Set = [EIei])
        Concatenate
         Set(Set = [Ll])
         Set(Set = [IYiy])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Uu])
       Set(Set = [Gg])
       Set(Set = [Uu])
       Set(Set = [Ssſ])
       Set(Set = [Tt])
      Concatenate
       Set(Set = [Ssſ])
       Set(Set = [Ee])
       Set(Set = [Pp])
       Set(Set = [Tt])
       Set(Set = [Ee])
       Set(Set = [Mm])
       Set(Set = [Bb])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Oo])
       Set(Set = [CKckK])
       Set(Set = [Tt])
       Set(Set = [Oo])
       Set(Set = [Bb])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Nn])
       Set(Set = [Oo])
       Set(Set = [Vv])
       Set(Set = [Ee])
       Set(Set = [Mm])
       Set(Set = [Bb])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Dd])
       Set(Set = [Ee])
       Set(Set = [CSZcszſ])
       Set(Set = [Ee])
       Set(Set = [Mm])
       Set(Set = [Bb])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Aa])
       Set(Set = [Nn])
      Concatenate
       Set(Set = [Ff])
       Set(Set = [Ee])
       Set(Set = [Bb])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [AaÄä])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Pp])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Uu])
       Set(Set = [LNln])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Uu])
       Set(Set = [Gg])
      Concatenate
       Set(Set = [Ssſ])
       Set(Set = [Ee])
       Set(Set = [Pp])
      Concatenate
       Set(Set = [Oo])
       Set(Set = [CKckK])
       Set(Set = [Tt])
      Concatenate
       Set(Set = [Nn])
       Set(Set = [Oo])
       Set(Set = [Vv])
      Concatenate
       Set(Set = [Dd])
       Set(Set = [Ee])
       Set(Set = [CZcz])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Aa])
       Set(Set = [Nn])
       Set(Set = [Uu])
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Ii])
      Concatenate
       Set(Set = [Ff])
       Set(Set = [Ee])
       Set(Set = [Bb])
       Set(Set = [Rr])
       Set(Set = [Uu])
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Ii])
      Concatenate
       Set(Set = [Mm])
       Alternate
        Concatenate
         Set(Set = [Aa])
         Set(Set = [Rr])
         Set(Set = [Ee])
         Set(Set = [Tt])
        Concatenate
         Set(Set = [Ee])
         Set(Set = [Ii])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Gg])
       Set(Set = [Uu])
       Set(Set = [Ssſ])
       Set(Set = [Tt])
       Set(Set = [Uu])
       Set(Set = [Ssſ])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Ää])
       SetloopAtomic(Set = [Nn])(Min = 2, Max = 2)
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Ff])
       Set(Set = [Ee])
       Set(Set = [Bb])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Ää])
       Set(Set = [Rr])
       Set(Set = [Zz])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Aa])
       Set(Set = [Nn])
       Set(Set = [Vv])
       Set(Set = [Ii])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Ff])
       Set(Set = [Éé])
       Set(Set = [Vv])
       Set(Set = [Rr])
       Set(Set = [Ii])
       Set(Set = [Ee])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Ssſ])
      Concatenate
       Set(Set = [Jj])
       Set(Set = [Uu])
       Set(Set = [Ii])
       Alternate
        Set(Set = [Nn])
        Concatenate
         SetloopAtomic(Set = [Ll])(Min = 2, Max = 2)
         Set(Set = [Ee])
         Set(Set = [Tt])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Oo])
       Set(Set = [Uu])
       Set(Set = [Tt])
      Concatenate
       Set(Set = [Ssſ])
       Set(Set = [Ee])
       Set(Set = [Pp])
       Set(Set = [Tt])
       Set(Set = [Ee])
       Set(Set = [Mm])
       Set(Set = [Bb])
       Set(Set = [Rr])
       Set(Set = [Ee])
      Concatenate
       Set(Set = [Oo])
       Set(Set = [Cc])
       Set(Set = [Tt])
       Set(Set = [Oo])
       Set(Set = [Bb])
       Set(Set = [Rr])
       Set(Set = [Ee])
      Concatenate
       Set(Set = [Nn])
       Set(Set = [Oo])
       Set(Set = [Vv])
       Set(Set = [Ee])
       Set(Set = [Mm])
       Set(Set = [Bb])
       Set(Set = [Rr])
       Set(Set = [Ee])
      Concatenate
       Set(Set = [Dd])
       Set(Set = [Éé])
       Set(Set = [Cc])
       Set(Set = [Ee])
       Set(Set = [Mm])
       Set(Set = [Bb])
       Set(Set = [Rr])
       Set(Set = [Ee])
      Concatenate
       Set(Set = [Oo])
       Set(Set = [Cc])
       Set(Set = [Aa])
       Set(Set = [KkK])
      Concatenate
       Set(Set = [Şş])
       Set(Set = [Uu])
       Set(Set = [Bb])
       Set(Set = [Aa])
       Set(Set = [Tt])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Tt])
      Concatenate
       Set(Set = [Nn])
       Set(Set = [Ii])
       Set(Set = [Ssſ])
       Set(Set = [Aa])
       Set(Set = [Nn])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Aa])
       Set(Set = [Yy])
       One(Ch = ı)
       Set(Set = [Ssſ])
      Concatenate
       Set(Set = [Hh])
       Set(Set = [Aa])
       Set(Set = [Zz])
       Set(Set = [Ii])
       Set(Set = [Rr])
       Set(Set = [Aa])
       Set(Set = [Nn])
      Concatenate
       Set(Set = [Tt])
       Set(Set = [Ee])
       SetloopAtomic(Set = [Mm])(Min = 2, Max = 2)
       Set(Set = [Uu])
       Set(Set = [Zz])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Ğğ])
       Set(Set = [Uu])
       Set(Set = [Ssſ])
       Set(Set = [Tt])
       Set(Set = [Oo])
       Set(Set = [Ssſ])
      Concatenate
       Set(Set = [Ee])
       Alternate
        Concatenate
         Set(Set = [Yy])
         Set(Set = [Ll])
         Set(Set = [Üü])
         Set(Set = [Ll])
        Concatenate
         Set(Set = [KkK])
         Set(Set = [Ii])
         Set(Set = [Mm])
      Concatenate
       Set(Set = [KkK])
       Set(Set = [Aa])
       Set(Set = [Ssſ])
       One(Ch = ı)
       Set(Set = [Mm])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Aa])
       Set(Set = [Ll])
       One(Ch = ı)
       Set(Set = [KkK])
      Concatenate
       Set(Set = [Oo])
       Set(Set = [Cc])
       Set(Set = [Aa])
      Concatenate
       Set(Set = [Şş])
       Set(Set = [Uu])
       Set(Set = [Bb])
      Concatenate
       Set(Set = [Mm])
       Set(Set = [Aa])
       Set(Set = [Rr])
      Concatenate
       Set(Set = [Nn])
       Set(Set = [Ii])
       Set(Set = [Ssſ])
      Concatenate
       Set(Set = [Hh])
       Set(Set = [Aa])
       Set(Set = [Zz])
      Concatenate
       Set(Set = [Tt])
       Set(Set = [Ee])
       Set(Set = [Mm])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Ğğ])
       Set(Set = [Uu])
      Concatenate
       Set(Set = [Ee])
       Alternate
        Concatenate
         Set(Set = [Yy])
         Set(Set = [Ll])
        Concatenate
         Set(Set = [KkK])
         Set(Set = [Ii])
      Concatenate
       Set(Set = [KkK])
       Set(Set = [Aa])
       Set(Set = [Ssſ])
      Concatenate
       Set(Set = [Aa])
       Set(Set = [Rr])
       Set(Set = [Aa])
    SetloopAtomic(Set = [,\.])(Min = 0, Max = 1)
    Set(Set = [\s])
    Capture(index = 6, unindex = -1)
     Atomic
      Alternate
       Concatenate
        Multi(String = "199")
        Set(Set = [0-9])
       Concatenate
        Multi(String = "20")
        Set(Set = [0-3])
        Set(Set = [0-9])
*/
// From main.go:17:34
// Pattern: "(?i)(January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|Maret|Mei|Agustus|Jänner|Feber|März|janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara)\\s([0-3]?[0-9])(?:st|nd|rd|th)?,?\\s(199[0-9]|20[0-3][0-9])|([0-3]?[0-9])(?:st|nd|rd|th|\\.)?\\s(?:of\\s)?(January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|Maret|Mei|Agustus|Jänner|Feber|März|janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara)[,.]?\\s(199[0-9]|20[0-3][0-9])"
// Options: regexp2.None
type rxLongDate_Engine struct{}

func (rxLongDate_Engine) Caps() map[int]int        { return nil }
func (rxLongDate_Engine) CapNames() map[string]int { return nil }
func (rxLongDate_Engine) CapsList() []string       { return nil }
func (rxLongDate_Engine) CapSize() int             { return 7 }

func (rxLongDate_Engine) FindFirstChar(r *regexp2.Runner) bool {
	pos := r.Runtextpos
	// Any possible match is at least 10 characters
	if pos <= len(r.Runtext)-10 {
		// The pattern begins with [0-9AD-FHJKM-OSTad-fhjkm-ostŞşſK]
		// Find the next occurrence. If it can't be found, there's no match.
		i := sNonAscii2425845632ac04d3dc6b5f60352d44d25fa35aaf888d12722b964c9ace2b4cd4.IndexOfAny(r.Runtext[pos:])
		if i >= 0 {
			r.Runtextpos = pos + i
			return true
		}

	}

	// No match found
	r.Runtextpos = len(r.Runtext)
	return false
}

func (rxLongDate_Engine) Execute(r *regexp2.Runner) error {
	atomic_stackpos := 0
	alternation_starting_pos := 0
	alternation_starting_capturepos := 0
	capture_starting_pos := 0
	alternation_starting_pos1 := 0
	alternation_starting_capturepos1 := 0
	alternation_branch := 0
	capture_starting_pos1 := 0
	var charloop_starting_pos, charloop_ending_pos = 0, 0
	charloop_capture_pos := 0
	loop_iteration := 0
	capture_starting_pos2 := 0
	capture_starting_pos3 := 0
	var charloop_starting_pos1, charloop_ending_pos1 = 0, 0
	charloop_capture_pos1 := 0
	loop_iteration1 := 0
	loop_iteration2 := 0
	capture_starting_pos4 := 0
	alternation_starting_pos2 := 0
	alternation_starting_capturepos2 := 0
	alternation_branch1 := 0
	var charloop_starting_pos2, charloop_ending_pos2 = 0, 0
	charloop_capture_pos2 := 0
	var charloop_starting_pos3, charloop_ending_pos3 = 0, 0
	charloop_capture_pos3 := 0
	capture_starting_pos5 := 0
	pos := r.Runtextpos
	matchStart := pos

	var slice = r.Runtext[pos:]

	// Node: Atomic
	// Atomic group.
	atomic_stackpos = r.Runstackpos

	// Node: Alternate
	// Match with 2 alternative expressions, atomically.
	alternation_starting_pos = pos
	alternation_starting_capturepos = r.Crawlpos()

	// Branch 0
	// Node: Concatenate
	// Node: Capture(index = 1, unindex = -1)
	// "1" capture group
	capture_starting_pos = pos

	// Node: Alternate
	// Match with 58 alternative expressions.
	alternation_starting_pos1 = pos
	alternation_starting_capturepos1 = r.Crawlpos()

	// Branch 0
	// Node: Concatenate
	if len(slice) < 6 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("januar")) /* Match the string "januar" (case-insensitive) */ {
		goto AlternationBranch1
	}

	// Node: SetloopAtomic(Set = [Yy])(Min = 0, Max = 1)
	// Match [Yy] atomically, optionally.
	if len(slice) > 6 && (slice[6]|0x20 == 'y') {
		slice = slice[1:]
		pos++
	}

	alternation_branch = 0
	pos += 6
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch1:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 1
	// Node: Concatenate
	if len(slice) < 7 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("februar")) /* Match the string "februar" (case-insensitive) */ {
		goto AlternationBranch2
	}

	// Node: SetloopAtomic(Set = [Yy])(Min = 0, Max = 1)
	// Match [Yy] atomically, optionally.
	if len(slice) > 7 && (slice[7]|0x20 == 'y') {
		slice = slice[1:]
		pos++
	}

	alternation_branch = 1
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch2:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 2
	// Node: Concatenate
	if len(slice) < 5 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("march")) /* Match the string "march" (case-insensitive) */ {
		goto AlternationBranch3
	}

	alternation_branch = 2
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch3:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 3
	// Node: Concatenate
	if len(slice) < 5 ||
		(slice[0]|0x20 != 'a') || /* Match [Aa]. */
		!helpers.IsInMask64(slice[1]-'P', 0x8200000082000000) || /* Match [PVpv]. */
		!helpers.StartsWithIgnoreCase(slice[2:], []rune("ril")) /* Match the string "ril" (case-insensitive) */ {
		goto AlternationBranch4
	}

	alternation_branch = 3
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch4:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 4
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("ma")) || /* Match the string "ma" (case-insensitive) */
		!helpers.IsInMask64(slice[2]-'I', 0x8000800080008000) /* Match [IYiy]. */ {
		goto AlternationBranch5
	}

	alternation_branch = 4
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch5:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 5
	// Node: Concatenate
	if len(slice) < 2 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("ju")) /* Match the string "ju" (case-insensitive) */ {
		goto AlternationBranch6
	}

	// Node: Alternate
	// Match with 2 alternative expressions.
	if len(slice) < 3 {
		goto AlternationBranch6
	}

	switch slice[2] {
	case 'N', 'n':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [EIei])
		// Match [EIei].
		if len(slice) < 4 || !helpers.IsInMask64(slice[3]-'E', 0x8800000088000000) {
			goto AlternationBranch6
		}

		pos += 4
		slice = r.Runtext[pos:]

	case 'L', 'l':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [IYiy])
		// Match [IYiy].
		if len(slice) < 4 || !helpers.IsInMask64(slice[3]-'I', 0x8000800080008000) {
			goto AlternationBranch6
		}

		pos += 4
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch6
	}

	alternation_branch = 5
	goto AlternationMatch1

AlternationBranch6:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 6
	// Node: Concatenate
	if len(slice) < 6 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("augu")) || /* Match the string "augu" (case-insensitive) */
		((slice[4]|0x20 != 's') && (slice[4] != 'ſ')) || /* Match [Ssſ]. */
		(slice[5]|0x20 != 't') /* Match [Tt]. */ {
		goto AlternationBranch7
	}

	alternation_branch = 6
	pos += 6
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch7:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 7
	// Node: Concatenate
	if len(slice) < 9 ||
		((slice[0]|0x20 != 's') && (slice[0] != 'ſ')) || /* Match [Ssſ]. */
		!helpers.StartsWithIgnoreCase(slice[1:], []rune("eptember")) /* Match the string "eptember" (case-insensitive) */ {
		goto AlternationBranch8
	}

	alternation_branch = 7
	pos += 9
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch8:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 8
	// Node: Concatenate
	if len(slice) < 7 ||
		(slice[0]|0x20 != 'o') || /* Match [Oo]. */
		!set_c585570961bd1835f850d6a9b56e3766bcb037a62b1dbc82686b27dd542a929b.CharIn(slice[1]) || /* Match [CKckK]. */
		!helpers.StartsWithIgnoreCase(slice[2:], []rune("tober")) /* Match the string "tober" (case-insensitive) */ {
		goto AlternationBranch9
	}

	alternation_branch = 8
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch9:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 9
	// Node: Concatenate
	if len(slice) < 8 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("november")) /* Match the string "november" (case-insensitive) */ {
		goto AlternationBranch10
	}

	alternation_branch = 9
	pos += 8
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch10:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 10
	// Node: Concatenate
	if len(slice) < 8 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("de")) || /* Match the string "de" (case-insensitive) */
		!set_54fec145c186052ca441a68ab7e5bcb6ce2ab384ff76b5756c3e7264dde6df4b.CharIn(slice[2]) || /* Match [CSZcszſ]. */
		!helpers.StartsWithIgnoreCase(slice[3:], []rune("ember")) /* Match the string "ember" (case-insensitive) */ {
		goto AlternationBranch11
	}

	alternation_branch = 10
	pos += 8
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch11:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 11
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("jan")) /* Match the string "jan" (case-insensitive) */ {
		goto AlternationBranch12
	}

	alternation_branch = 11
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch12:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 12
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("feb")) /* Match the string "feb" (case-insensitive) */ {
		goto AlternationBranch13
	}

	alternation_branch = 12
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch13:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 13
	// Node: Concatenate
	if len(slice) < 3 ||
		(slice[0]|0x20 != 'm') || /* Match [Mm]. */
		!set_0aed9828638bbe7e52933e044a08751d611ecbcf65fa3d7b353e7a1eecc9c411.CharIn(slice[1]) || /* Match [AaÄä]. */
		(slice[2]|0x20 != 'r') /* Match [Rr]. */ {
		goto AlternationBranch14
	}

	alternation_branch = 13
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch14:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 14
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("apr")) /* Match the string "apr" (case-insensitive) */ {
		goto AlternationBranch15
	}

	alternation_branch = 14
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch15:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 15
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("ju")) || /* Match the string "ju" (case-insensitive) */
		!helpers.IsInMask64(slice[2]-'L', 0xa0000000a0000000) /* Match [LNln]. */ {
		goto AlternationBranch16
	}

	alternation_branch = 15
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch16:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 16
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("aug")) /* Match the string "aug" (case-insensitive) */ {
		goto AlternationBranch17
	}

	alternation_branch = 16
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch17:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 17
	// Node: Concatenate
	if len(slice) < 3 ||
		((slice[0]|0x20 != 's') && (slice[0] != 'ſ')) || /* Match [Ssſ]. */
		!helpers.StartsWithIgnoreCase(slice[1:], []rune("ep")) /* Match the string "ep" (case-insensitive) */ {
		goto AlternationBranch18
	}

	alternation_branch = 17
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch18:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 18
	// Node: Concatenate
	if len(slice) < 3 ||
		(slice[0]|0x20 != 'o') || /* Match [Oo]. */
		!set_c585570961bd1835f850d6a9b56e3766bcb037a62b1dbc82686b27dd542a929b.CharIn(slice[1]) || /* Match [CKckK]. */
		(slice[2]|0x20 != 't') /* Match [Tt]. */ {
		goto AlternationBranch19
	}

	alternation_branch = 18
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch19:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 19
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("nov")) /* Match the string "nov" (case-insensitive) */ {
		goto AlternationBranch20
	}

	alternation_branch = 19
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch20:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 20
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("de")) || /* Match the string "de" (case-insensitive) */
		!helpers.IsInMask64(slice[2]-'C', 0x8000010080000100) /* Match [CZcz]. */ {
		goto AlternationBranch21
	}

	alternation_branch = 20
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch21:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 21
	// Node: Concatenate
	if len(slice) < 7 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("januari")) /* Match the string "januari" (case-insensitive) */ {
		goto AlternationBranch22
	}

	alternation_branch = 21
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch22:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 22
	// Node: Concatenate
	if len(slice) < 8 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("februari")) /* Match the string "februari" (case-insensitive) */ {
		goto AlternationBranch23
	}

	alternation_branch = 22
	pos += 8
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch23:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 23
	// Node: Concatenate
	// Node: Set(Set = [Mm])
	// Match [Mm].
	if len(slice) == 0 || (slice[0]|0x20 != 'm') {
		goto AlternationBranch24
	}

	// Node: Alternate
	// Match with 2 alternative expressions.
	if len(slice) < 2 {
		goto AlternationBranch24
	}

	switch slice[1] {
	case 'A', 'a':
		// Node: Concatenate
		// Node: Empty

		if len(slice) < 5 ||
			!helpers.StartsWithIgnoreCase(slice[2:], []rune("ret")) /* Match the string "ret" (case-insensitive) */ {
			goto AlternationBranch24
		}

		pos += 5
		slice = r.Runtext[pos:]

	case 'E', 'e':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Ii])
		// Match [Ii].
		if len(slice) < 3 || (slice[2]|0x20 != 'i') {
			goto AlternationBranch24
		}

		pos += 3
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch24
	}

	alternation_branch = 23
	goto AlternationMatch1

AlternationBranch24:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 24
	// Node: Concatenate
	if len(slice) < 7 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("agu")) || /* Match the string "agu" (case-insensitive) */
		((slice[3]|0x20 != 's') && (slice[3] != 'ſ')) || /* Match [Ssſ]. */
		!helpers.StartsWithIgnoreCase(slice[4:], []rune("tu")) || /* Match the string "tu" (case-insensitive) */
		((slice[6]|0x20 != 's') && (slice[6] != 'ſ')) /* Match [Ssſ]. */ {
		goto AlternationBranch25
	}

	alternation_branch = 24
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch25:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 25
	// Node: Concatenate
	if len(slice) < 6 ||
		(slice[0]|0x20 != 'j') || /* Match [Jj]. */
		(slice[1]|0x20 != 'ä') || /* Match [Ää]. */
		!helpers.StartsWithIgnoreCase(slice[2:], []rune("nner")) /* Match the string "nner" (case-insensitive) */ {
		goto AlternationBranch26
	}

	alternation_branch = 25
	pos += 6
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch26:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 26
	// Node: Concatenate
	if len(slice) < 5 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("feber")) /* Match the string "feber" (case-insensitive) */ {
		goto AlternationBranch27
	}

	alternation_branch = 26
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch27:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 27
	// Node: Concatenate
	if len(slice) < 4 ||
		(slice[0]|0x20 != 'm') || /* Match [Mm]. */
		(slice[1]|0x20 != 'ä') || /* Match [Ää]. */
		!helpers.StartsWithIgnoreCase(slice[2:], []rune("rz")) /* Match the string "rz" (case-insensitive) */ {
		goto AlternationBranch28
	}

	alternation_branch = 27
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch28:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 28
	// Node: Concatenate
	if len(slice) < 7 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("janvier")) /* Match the string "janvier" (case-insensitive) */ {
		goto AlternationBranch29
	}

	alternation_branch = 28
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch29:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 29
	// Node: Concatenate
	if len(slice) < 7 ||
		(slice[0]|0x20 != 'f') || /* Match [Ff]. */
		(slice[1]|0x20 != 'é') || /* Match [Éé]. */
		!helpers.StartsWithIgnoreCase(slice[2:], []rune("vrier")) /* Match the string "vrier" (case-insensitive) */ {
		goto AlternationBranch30
	}

	alternation_branch = 29
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch30:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 30
	// Node: Concatenate
	if len(slice) < 4 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("mar")) || /* Match the string "mar" (case-insensitive) */
		((slice[3]|0x20 != 's') && (slice[3] != 'ſ')) /* Match [Ssſ]. */ {
		goto AlternationBranch31
	}

	alternation_branch = 30
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch31:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 31
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("jui")) /* Match the string "jui" (case-insensitive) */ {
		goto AlternationBranch32
	}

	// Node: Alternate
	// Match with 2 alternative expressions.
	if len(slice) < 4 {
		goto AlternationBranch32
	}

	switch slice[3] {
	case 'N', 'n':
		pos += 4
		slice = r.Runtext[pos:]

	case 'L', 'l':
		// Node: Concatenate
		if len(slice) < 7 ||
			!helpers.StartsWithIgnoreCase(slice[3:], []rune("llet")) /* Match the string "llet" (case-insensitive) */ {
			goto AlternationBranch32
		}

		pos += 7
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch32
	}

	alternation_branch = 31
	goto AlternationMatch1

AlternationBranch32:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 32
	// Node: Concatenate
	if len(slice) < 4 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("aout")) /* Match the string "aout" (case-insensitive) */ {
		goto AlternationBranch33
	}

	alternation_branch = 32
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch33:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 33
	// Node: Concatenate
	if len(slice) < 9 ||
		((slice[0]|0x20 != 's') && (slice[0] != 'ſ')) || /* Match [Ssſ]. */
		!helpers.StartsWithIgnoreCase(slice[1:], []rune("eptembre")) /* Match the string "eptembre" (case-insensitive) */ {
		goto AlternationBranch34
	}

	alternation_branch = 33
	pos += 9
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch34:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 34
	// Node: Concatenate
	if len(slice) < 7 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("octobre")) /* Match the string "octobre" (case-insensitive) */ {
		goto AlternationBranch35
	}

	alternation_branch = 34
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch35:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 35
	// Node: Concatenate
	if len(slice) < 8 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("novembre")) /* Match the string "novembre" (case-insensitive) */ {
		goto AlternationBranch36
	}

	alternation_branch = 35
	pos += 8
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch36:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 36
	// Node: Concatenate
	if len(slice) < 8 ||
		(slice[0]|0x20 != 'd') || /* Match [Dd]. */
		(slice[1]|0x20 != 'é') || /* Match [Éé]. */
		!helpers.StartsWithIgnoreCase(slice[2:], []rune("cembre")) /* Match the string "cembre" (case-insensitive) */ {
		goto AlternationBranch37
	}

	alternation_branch = 36
	pos += 8
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch37:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 37
	// Node: Concatenate
	if len(slice) < 4 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("oca")) || /* Match the string "oca" (case-insensitive) */
		((slice[3]|0x20 != 'k') && (slice[3] != 'K')) /* Match [KkK]. */ {
		goto AlternationBranch38
	}

	alternation_branch = 37
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch38:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 38
	// Node: Concatenate
	if len(slice) < 5 ||
		!helpers.IsBetween(slice[0], 'Ş', 'ş') || /* Match [Şş]. */
		!helpers.StartsWithIgnoreCase(slice[1:], []rune("ubat")) /* Match the string "ubat" (case-insensitive) */ {
		goto AlternationBranch39
	}

	alternation_branch = 38
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch39:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 39
	// Node: Concatenate
	if len(slice) < 4 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("mart")) /* Match the string "mart" (case-insensitive) */ {
		goto AlternationBranch40
	}

	alternation_branch = 39
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch40:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 40
	// Node: Concatenate
	if len(slice) < 5 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("ni")) || /* Match the string "ni" (case-insensitive) */
		((slice[2]|0x20 != 's') && (slice[2] != 'ſ')) || /* Match [Ssſ]. */
		!helpers.StartsWithIgnoreCase(slice[3:], []rune("an")) /* Match the string "an" (case-insensitive) */ {
		goto AlternationBranch41
	}

	alternation_branch = 40
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch41:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 41
	// Node: Concatenate
	if len(slice) < 5 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("may")) || /* Match the string "may" (case-insensitive) */
		slice[3] != 'ı' || /* Match 'ı'. */
		((slice[4]|0x20 != 's') && (slice[4] != 'ſ')) /* Match [Ssſ]. */ {
		goto AlternationBranch42
	}

	alternation_branch = 41
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch42:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 42
	// Node: Concatenate
	if len(slice) < 7 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("haziran")) /* Match the string "haziran" (case-insensitive) */ {
		goto AlternationBranch43
	}

	alternation_branch = 42
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch43:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 43
	// Node: Concatenate
	if len(slice) < 6 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("temmuz")) /* Match the string "temmuz" (case-insensitive) */ {
		goto AlternationBranch44
	}

	alternation_branch = 43
	pos += 6
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch44:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 44
	// Node: Concatenate
	if len(slice) < 7 ||
		(slice[0]|0x20 != 'a') || /* Match [Aa]. */
		!helpers.IsBetween(slice[1], 'Ğ', 'ğ') || /* Match [Ğğ]. */
		(slice[2]|0x20 != 'u') || /* Match [Uu]. */
		((slice[3]|0x20 != 's') && (slice[3] != 'ſ')) || /* Match [Ssſ]. */
		!helpers.StartsWithIgnoreCase(slice[4:], []rune("to")) || /* Match the string "to" (case-insensitive) */
		((slice[6]|0x20 != 's') && (slice[6] != 'ſ')) /* Match [Ssſ]. */ {
		goto AlternationBranch45
	}

	alternation_branch = 44
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch45:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 45
	// Node: Concatenate
	// Node: Set(Set = [Ee])
	// Match [Ee].
	if len(slice) == 0 || (slice[0]|0x20 != 'e') {
		goto AlternationBranch46
	}

	// Node: Alternate
	// Match with 2 alternative expressions.
	if len(slice) < 2 {
		goto AlternationBranch46
	}

	switch slice[1] {
	case 'Y', 'y':
		// Node: Concatenate
		// Node: Empty

		if len(slice) < 5 ||
			(slice[2]|0x20 != 'l') || /* Match [Ll]. */
			(slice[3]|0x20 != 'ü') || /* Match [Üü]. */
			(slice[4]|0x20 != 'l') /* Match [Ll]. */ {
			goto AlternationBranch46
		}

		pos += 5
		slice = r.Runtext[pos:]

	case 'K', 'k', 'K':
		// Node: Concatenate
		// Node: Empty

		if len(slice) < 4 ||
			!helpers.StartsWithIgnoreCase(slice[2:], []rune("im")) /* Match the string "im" (case-insensitive) */ {
			goto AlternationBranch46
		}

		pos += 4
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch46
	}

	alternation_branch = 45
	goto AlternationMatch1

AlternationBranch46:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 46
	// Node: Concatenate
	if len(slice) < 5 ||
		((slice[0]|0x20 != 'k') && (slice[0] != 'K')) || /* Match [KkK]. */
		(slice[1]|0x20 != 'a') || /* Match [Aa]. */
		((slice[2]|0x20 != 's') && (slice[2] != 'ſ')) || /* Match [Ssſ]. */
		slice[3] != 'ı' || /* Match 'ı'. */
		(slice[4]|0x20 != 'm') /* Match [Mm]. */ {
		goto AlternationBranch47
	}

	alternation_branch = 46
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch47:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 47
	// Node: Concatenate
	if len(slice) < 6 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("aral")) || /* Match the string "aral" (case-insensitive) */
		slice[4] != 'ı' || /* Match 'ı'. */
		((slice[5]|0x20 != 'k') && (slice[5] != 'K')) /* Match [KkK]. */ {
		goto AlternationBranch48
	}

	alternation_branch = 47
	pos += 6
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch48:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 48
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("oca")) /* Match the string "oca" (case-insensitive) */ {
		goto AlternationBranch49
	}

	alternation_branch = 48
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch49:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 49
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.IsBetween(slice[0], 'Ş', 'ş') || /* Match [Şş]. */
		!helpers.StartsWithIgnoreCase(slice[1:], []rune("ub")) /* Match the string "ub" (case-insensitive) */ {
		goto AlternationBranch50
	}

	alternation_branch = 49
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch50:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 50
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("mar")) /* Match the string "mar" (case-insensitive) */ {
		goto AlternationBranch51
	}

	alternation_branch = 50
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch51:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 51
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("ni")) || /* Match the string "ni" (case-insensitive) */
		((slice[2]|0x20 != 's') && (slice[2] != 'ſ')) /* Match [Ssſ]. */ {
		goto AlternationBranch52
	}

	alternation_branch = 51
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch52:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 52
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("haz")) /* Match the string "haz" (case-insensitive) */ {
		goto AlternationBranch53
	}

	alternation_branch = 52
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch53:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 53
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("tem")) /* Match the string "tem" (case-insensitive) */ {
		goto AlternationBranch54
	}

	alternation_branch = 53
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch54:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 54
	// Node: Concatenate
	if len(slice) < 3 ||
		(slice[0]|0x20 != 'a') || /* Match [Aa]. */
		!helpers.IsBetween(slice[1], 'Ğ', 'ğ') || /* Match [Ğğ]. */
		(slice[2]|0x20 != 'u') /* Match [Uu]. */ {
		goto AlternationBranch55
	}

	alternation_branch = 54
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch55:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 55
	// Node: Concatenate
	// Node: Set(Set = [Ee])
	// Match [Ee].
	if len(slice) == 0 || (slice[0]|0x20 != 'e') {
		goto AlternationBranch56
	}

	// Node: Alternate
	// Match with 2 alternative expressions.
	if len(slice) < 2 {
		goto AlternationBranch56
	}

	switch slice[1] {
	case 'Y', 'y':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Ll])
		// Match [Ll].
		if len(slice) < 3 || (slice[2]|0x20 != 'l') {
			goto AlternationBranch56
		}

		pos += 3
		slice = r.Runtext[pos:]

	case 'K', 'k', 'K':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Ii])
		// Match [Ii].
		if len(slice) < 3 || (slice[2]|0x20 != 'i') {
			goto AlternationBranch56
		}

		pos += 3
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch56
	}

	alternation_branch = 55
	goto AlternationMatch1

AlternationBranch56:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 56
	// Node: Concatenate
	if len(slice) < 3 ||
		((slice[0]|0x20 != 'k') && (slice[0] != 'K')) || /* Match [KkK]. */
		(slice[1]|0x20 != 'a') || /* Match [Aa]. */
		((slice[2]|0x20 != 's') && (slice[2] != 'ſ')) /* Match [Ssſ]. */ {
		goto AlternationBranch57
	}

	alternation_branch = 56
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBranch57:
	pos = alternation_starting_pos1
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos1)

	// Branch 57
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("ara")) /* Match the string "ara" (case-insensitive) */ {
		goto AlternationBranch
	}

	alternation_branch = 57
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch1

AlternationBacktrack1:
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	switch alternation_branch {
	case 0:
		goto AlternationBranch1
	case 1:
		goto AlternationBranch2
	case 2:
		goto AlternationBranch3
	case 3:
		goto AlternationBranch4
	case 4:
		goto AlternationBranch5
	case 5:
		goto AlternationBranch6
	case 6:
		goto AlternationBranch7
	case 7:
		goto AlternationBranch8
	case 8:
		goto AlternationBranch9
	case 9:
		goto AlternationBranch10
	case 10:
		goto AlternationBranch11
	case 11:
		goto AlternationBranch12
	case 12:
		goto AlternationBranch13
	case 13:
		goto AlternationBranch14
	case 14:
		goto AlternationBranch15
	case 15:
		goto AlternationBranch16
	case 16:
		goto AlternationBranch17
	case 17:
		goto AlternationBranch18
	case 18:
		goto AlternationBranch19
	case 19:
		goto AlternationBranch20
	case 20:
		goto AlternationBranch21
	case 21:
		goto AlternationBranch22
	case 22:
		goto AlternationBranch23
	case 23:
		goto AlternationBranch24
	case 24:
		goto AlternationBranch25
	case 25:
		goto AlternationBranch26
	case 26:
		goto AlternationBranch27
	case 27:
		goto AlternationBranch28
	case 28:
		goto AlternationBranch29
	case 29:
		goto AlternationBranch30
	case 30:
		goto AlternationBranch31
	case 31:
		goto AlternationBranch32
	case 32:
		goto AlternationBranch33
	case 33:
		goto AlternationBranch34
	case 34:
		goto AlternationBranch35
	case 35:
		goto AlternationBranch36
	case 36:
		goto AlternationBranch37
	case 37:
		goto AlternationBranch38
	case 38:
		goto AlternationBranch39
	case 39:
		goto AlternationBranch40
	case 40:
		goto AlternationBranch41
	case 41:
		goto AlternationBranch42
	case 42:
		goto AlternationBranch43
	case 43:
		goto AlternationBranch44
	case 44:
		goto AlternationBranch45
	case 45:
		goto AlternationBranch46
	case 46:
		goto AlternationBranch47
	case 47:
		goto AlternationBranch48
	case 48:
		goto AlternationBranch49
	case 49:
		goto AlternationBranch50
	case 50:
		goto AlternationBranch51
	case 51:
		goto AlternationBranch52
	case 52:
		goto AlternationBranch53
	case 53:
		goto AlternationBranch54
	case 54:
		goto AlternationBranch55
	case 55:
		goto AlternationBranch56
	case 56:
		goto AlternationBranch57
	case 57:
		goto AlternationBranch
	}

AlternationMatch1:
	;

	r.Capture(1, capture_starting_pos, pos)

	goto CaptureSkipBacktrack

CaptureBacktrack:
	goto AlternationBacktrack1

CaptureSkipBacktrack:
	;

	// Node: Set(Set = [\s])
	// Match [\s].
	if len(slice) == 0 || !unicode.IsSpace(slice[0]) {
		goto CaptureBacktrack
	}

	// Node: Capture(index = 2, unindex = -1)
	// "2" capture group
	pos++
	slice = r.Runtext[pos:]
	capture_starting_pos1 = pos

	// Node: Concatenate
	// Node: Setloop(Set = [0-3])(Min = 0, Max = 1)
	// Match [0-3] greedily, optionally.
	charloop_starting_pos = pos

	if len(slice) > 0 && helpers.IsBetween(slice[0], '0', '3') {
		slice = slice[1:]
		pos++
	}

	charloop_ending_pos = pos
	goto CharLoopEnd

CharLoopBacktrack:
	r.UncaptureUntil(charloop_capture_pos)

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos >= charloop_ending_pos {
		goto CaptureBacktrack
	}
	charloop_ending_pos--
	pos = charloop_ending_pos
	slice = r.Runtext[pos:]

CharLoopEnd:
	charloop_capture_pos = r.Crawlpos()

	// Node: Set(Set = [0-9])
	// Match [0-9].
	if len(slice) == 0 || !helpers.IsBetween(slice[0], '0', '9') {
		goto CharLoopBacktrack
	}

	pos++
	slice = r.Runtext[pos:]
	r.Capture(2, capture_starting_pos1, pos)

	goto CaptureSkipBacktrack1

CaptureBacktrack1:
	goto CharLoopBacktrack

CaptureSkipBacktrack1:
	;

	// Node: Loop(Min = 0, Max = 1)
	// Optional (greedy).
	loop_iteration = 0

LoopBody:
	r.StackPush2(r.Crawlpos(), pos)

	loop_iteration++

	// Node: Alternate
	// Match with 4 alternative expressions.
	if len(slice) == 0 {
		goto LoopIterationNoMatch
	}

	switch slice[0] {
	case 'S', 's', 'ſ':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Tt])
		// Match [Tt].
		if len(slice) < 2 || (slice[1]|0x20 != 't') {
			goto LoopIterationNoMatch
		}

		pos += 2
		slice = r.Runtext[pos:]

	case 'N', 'n':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Dd])
		// Match [Dd].
		if len(slice) < 2 || (slice[1]|0x20 != 'd') {
			goto LoopIterationNoMatch
		}

		pos += 2
		slice = r.Runtext[pos:]

	case 'R', 'r':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Dd])
		// Match [Dd].
		if len(slice) < 2 || (slice[1]|0x20 != 'd') {
			goto LoopIterationNoMatch
		}

		pos += 2
		slice = r.Runtext[pos:]

	case 'T', 't':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Hh])
		// Match [Hh].
		if len(slice) < 2 || (slice[1]|0x20 != 'h') {
			goto LoopIterationNoMatch
		}

		pos += 2
		slice = r.Runtext[pos:]

	default:
		goto LoopIterationNoMatch
	}

	// The loop has an upper bound of 1. Continue iterating greedily if it hasn't yet been reached.
	if loop_iteration == 0 {
		goto LoopBody
	}
	goto LoopEnd

	// The loop iteration failed. Put state back to the way it was before the iteration.
LoopIterationNoMatch:
	loop_iteration--
	if loop_iteration < 0 {
		// Unable to match the remainder of the expression after exhausting the loop.
		goto CaptureBacktrack1
	}
	pos = r.StackPop()
	r.UncaptureUntil(r.StackPop())
	slice = r.Runtext[pos:]
LoopEnd:
	;

	// Node: OneloopAtomic(Ch = ,)(Min = 0, Max = 1)
	// Match ',' atomically, optionally.
	if len(slice) > 0 && slice[0] == ',' {
		slice = slice[1:]
		pos++
	}

	// Node: Set(Set = [\s])
	// Match [\s].
	if len(slice) == 0 || !unicode.IsSpace(slice[0]) {
		goto LoopIterationNoMatch
	}

	// Node: Capture(index = 3, unindex = -1)
	// "3" capture group
	pos++
	slice = r.Runtext[pos:]
	capture_starting_pos2 = pos

	// Node: Atomic
	// Node: Alternate
	// Match with 2 alternative expressions, atomically.
	if len(slice) == 0 {
		goto LoopIterationNoMatch
	}

	switch slice[0] {
	case '1':
		// Node: Multi(String = "99")
		// Match the string "99".
		if !helpers.StartsWith(slice[1:], []rune("99")) {
			goto LoopIterationNoMatch
		}

		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [0-9])
		// Match [0-9].
		if len(slice) < 4 || !helpers.IsBetween(slice[3], '0', '9') {
			goto LoopIterationNoMatch
		}

		pos += 4
		slice = r.Runtext[pos:]

	case '2':
		// Node: One(Ch = 0)
		// Match '0'.
		if len(slice) < 2 || slice[1] != '0' {
			goto LoopIterationNoMatch
		}

		// Node: Concatenate
		// Node: Empty

		if len(slice) < 4 ||
			!helpers.IsBetween(slice[2], '0', '3') || /* Match [0-3]. */
			!helpers.IsBetween(slice[3], '0', '9') /* Match [0-9]. */ {
			goto LoopIterationNoMatch
		}

		pos += 4
		slice = r.Runtext[pos:]

	default:
		goto LoopIterationNoMatch
	}

	r.Capture(3, capture_starting_pos2, pos)

	goto AlternationMatch

AlternationBranch:
	pos = alternation_starting_pos
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos)

	// Branch 1
	// Node: Concatenate
	// Node: Capture(index = 4, unindex = -1)
	// "4" capture group
	capture_starting_pos3 = pos

	// Node: Concatenate
	// Node: Setloop(Set = [0-3])(Min = 0, Max = 1)
	// Match [0-3] greedily, optionally.
	charloop_starting_pos1 = pos

	if len(slice) > 0 && helpers.IsBetween(slice[0], '0', '3') {
		slice = slice[1:]
		pos++
	}

	charloop_ending_pos1 = pos
	goto CharLoopEnd1

CharLoopBacktrack1:
	r.UncaptureUntil(charloop_capture_pos1)

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos1 >= charloop_ending_pos1 {
		r.UncaptureUntil(0)
		return nil // The input didn't match.
	}
	charloop_ending_pos1--
	pos = charloop_ending_pos1
	slice = r.Runtext[pos:]

CharLoopEnd1:
	charloop_capture_pos1 = r.Crawlpos()

	// Node: Set(Set = [0-9])
	// Match [0-9].
	if len(slice) == 0 || !helpers.IsBetween(slice[0], '0', '9') {
		goto CharLoopBacktrack1
	}

	pos++
	slice = r.Runtext[pos:]
	r.Capture(4, capture_starting_pos3, pos)

	goto CaptureSkipBacktrack2

CaptureBacktrack2:
	goto CharLoopBacktrack1

CaptureSkipBacktrack2:
	;

	// Node: Loop(Min = 0, Max = 1)
	// Optional (greedy).
	loop_iteration1 = 0

LoopBody1:
	r.StackPush2(r.Crawlpos(), pos)

	loop_iteration1++

	// Node: Alternate
	// Match with 5 alternative expressions.
	if len(slice) == 0 {
		goto LoopIterationNoMatch1
	}

	switch slice[0] {
	case 'S', 's', 'ſ':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Tt])
		// Match [Tt].
		if len(slice) < 2 || (slice[1]|0x20 != 't') {
			goto LoopIterationNoMatch1
		}

		pos += 2
		slice = r.Runtext[pos:]

	case 'N', 'n':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Dd])
		// Match [Dd].
		if len(slice) < 2 || (slice[1]|0x20 != 'd') {
			goto LoopIterationNoMatch1
		}

		pos += 2
		slice = r.Runtext[pos:]

	case 'R', 'r':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Dd])
		// Match [Dd].
		if len(slice) < 2 || (slice[1]|0x20 != 'd') {
			goto LoopIterationNoMatch1
		}

		pos += 2
		slice = r.Runtext[pos:]

	case 'T', 't':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Hh])
		// Match [Hh].
		if len(slice) < 2 || (slice[1]|0x20 != 'h') {
			goto LoopIterationNoMatch1
		}

		pos += 2
		slice = r.Runtext[pos:]

	case '.':
		pos++
		slice = r.Runtext[pos:]

	default:
		goto LoopIterationNoMatch1
	}

	// The loop has an upper bound of 1. Continue iterating greedily if it hasn't yet been reached.
	if loop_iteration1 == 0 {
		goto LoopBody1
	}
	goto LoopEnd1

	// The loop iteration failed. Put state back to the way it was before the iteration.
LoopIterationNoMatch1:
	loop_iteration1--
	if loop_iteration1 < 0 {
		// Unable to match the remainder of the expression after exhausting the loop.
		goto CaptureBacktrack2
	}
	pos = r.StackPop()
	r.UncaptureUntil(r.StackPop())
	slice = r.Runtext[pos:]
LoopEnd1:
	;

	// Node: Set(Set = [\s])
	// Match [\s].
	if len(slice) == 0 || !unicode.IsSpace(slice[0]) {
		goto LoopIterationNoMatch1
	}

	// Node: Loop(Min = 0, Max = 1)
	// Optional (greedy).
	pos++
	slice = r.Runtext[pos:]
	loop_iteration2 = 0

LoopBody2:
	r.StackPush2(r.Crawlpos(), pos)

	loop_iteration2++

	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("of")) || /* Match the string "of" (case-insensitive) */
		!unicode.IsSpace(slice[2]) /* Match [\s]. */ {
		goto LoopIterationNoMatch2
	}

	pos += 3
	slice = r.Runtext[pos:]

	// The loop has an upper bound of 1. Continue iterating greedily if it hasn't yet been reached.
	if loop_iteration2 == 0 {
		goto LoopBody2
	}
	goto LoopEnd2

	// The loop iteration failed. Put state back to the way it was before the iteration.
LoopIterationNoMatch2:
	loop_iteration2--
	if loop_iteration2 < 0 {
		// Unable to match the remainder of the expression after exhausting the loop.
		goto LoopIterationNoMatch1
	}
	pos = r.StackPop()
	r.UncaptureUntil(r.StackPop())
	slice = r.Runtext[pos:]
LoopEnd2:
	;

	// Node: Capture(index = 5, unindex = -1)
	// "5" capture group
	capture_starting_pos4 = pos

	// Node: Alternate
	// Match with 58 alternative expressions.
	alternation_starting_pos2 = pos
	alternation_starting_capturepos2 = r.Crawlpos()

	// Branch 0
	// Node: Concatenate
	if len(slice) < 6 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("januar")) /* Match the string "januar" (case-insensitive) */ {
		goto AlternationBranch58
	}

	// Node: Setloop(Set = [Yy])(Min = 0, Max = 1)
	// Match [Yy] greedily, optionally.
	pos += 6
	slice = r.Runtext[pos:]
	charloop_starting_pos2 = pos

	if len(slice) > 0 && (slice[0]|0x20 == 'y') {
		slice = slice[1:]
		pos++
	}

	charloop_ending_pos2 = pos
	goto CharLoopEnd2

CharLoopBacktrack2:
	r.UncaptureUntil(charloop_capture_pos2)

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos2 >= charloop_ending_pos2 {
		goto AlternationBranch58
	}
	charloop_ending_pos2--
	pos = charloop_ending_pos2
	slice = r.Runtext[pos:]

CharLoopEnd2:
	charloop_capture_pos2 = r.Crawlpos()

	alternation_branch1 = 0
	goto AlternationMatch2

AlternationBranch58:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 1
	// Node: Concatenate
	if len(slice) < 7 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("februar")) /* Match the string "februar" (case-insensitive) */ {
		goto AlternationBranch59
	}

	// Node: Setloop(Set = [Yy])(Min = 0, Max = 1)
	// Match [Yy] greedily, optionally.
	pos += 7
	slice = r.Runtext[pos:]
	charloop_starting_pos3 = pos

	if len(slice) > 0 && (slice[0]|0x20 == 'y') {
		slice = slice[1:]
		pos++
	}

	charloop_ending_pos3 = pos
	goto CharLoopEnd3

CharLoopBacktrack3:
	r.UncaptureUntil(charloop_capture_pos3)

	if err := r.CheckTimeout(); err != nil {
		return err
	}
	if charloop_starting_pos3 >= charloop_ending_pos3 {
		goto AlternationBranch59
	}
	charloop_ending_pos3--
	pos = charloop_ending_pos3
	slice = r.Runtext[pos:]

CharLoopEnd3:
	charloop_capture_pos3 = r.Crawlpos()

	alternation_branch1 = 1
	goto AlternationMatch2

AlternationBranch59:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 2
	// Node: Concatenate
	if len(slice) < 5 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("march")) /* Match the string "march" (case-insensitive) */ {
		goto AlternationBranch60
	}

	alternation_branch1 = 2
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch60:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 3
	// Node: Concatenate
	if len(slice) < 5 ||
		(slice[0]|0x20 != 'a') || /* Match [Aa]. */
		!helpers.IsInMask64(slice[1]-'P', 0x8200000082000000) || /* Match [PVpv]. */
		!helpers.StartsWithIgnoreCase(slice[2:], []rune("ril")) /* Match the string "ril" (case-insensitive) */ {
		goto AlternationBranch61
	}

	alternation_branch1 = 3
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch61:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 4
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("ma")) || /* Match the string "ma" (case-insensitive) */
		!helpers.IsInMask64(slice[2]-'I', 0x8000800080008000) /* Match [IYiy]. */ {
		goto AlternationBranch62
	}

	alternation_branch1 = 4
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch62:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 5
	// Node: Concatenate
	if len(slice) < 2 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("ju")) /* Match the string "ju" (case-insensitive) */ {
		goto AlternationBranch63
	}

	// Node: Alternate
	// Match with 2 alternative expressions.
	if len(slice) < 3 {
		goto AlternationBranch63
	}

	switch slice[2] {
	case 'N', 'n':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [EIei])
		// Match [EIei].
		if len(slice) < 4 || !helpers.IsInMask64(slice[3]-'E', 0x8800000088000000) {
			goto AlternationBranch63
		}

		pos += 4
		slice = r.Runtext[pos:]

	case 'L', 'l':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [IYiy])
		// Match [IYiy].
		if len(slice) < 4 || !helpers.IsInMask64(slice[3]-'I', 0x8000800080008000) {
			goto AlternationBranch63
		}

		pos += 4
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch63
	}

	alternation_branch1 = 5
	goto AlternationMatch2

AlternationBranch63:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 6
	// Node: Concatenate
	if len(slice) < 6 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("augu")) || /* Match the string "augu" (case-insensitive) */
		((slice[4]|0x20 != 's') && (slice[4] != 'ſ')) || /* Match [Ssſ]. */
		(slice[5]|0x20 != 't') /* Match [Tt]. */ {
		goto AlternationBranch64
	}

	alternation_branch1 = 6
	pos += 6
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch64:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 7
	// Node: Concatenate
	if len(slice) < 9 ||
		((slice[0]|0x20 != 's') && (slice[0] != 'ſ')) || /* Match [Ssſ]. */
		!helpers.StartsWithIgnoreCase(slice[1:], []rune("eptember")) /* Match the string "eptember" (case-insensitive) */ {
		goto AlternationBranch65
	}

	alternation_branch1 = 7
	pos += 9
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch65:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 8
	// Node: Concatenate
	if len(slice) < 7 ||
		(slice[0]|0x20 != 'o') || /* Match [Oo]. */
		!set_c585570961bd1835f850d6a9b56e3766bcb037a62b1dbc82686b27dd542a929b.CharIn(slice[1]) || /* Match [CKckK]. */
		!helpers.StartsWithIgnoreCase(slice[2:], []rune("tober")) /* Match the string "tober" (case-insensitive) */ {
		goto AlternationBranch66
	}

	alternation_branch1 = 8
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch66:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 9
	// Node: Concatenate
	if len(slice) < 8 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("november")) /* Match the string "november" (case-insensitive) */ {
		goto AlternationBranch67
	}

	alternation_branch1 = 9
	pos += 8
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch67:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 10
	// Node: Concatenate
	if len(slice) < 8 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("de")) || /* Match the string "de" (case-insensitive) */
		!set_54fec145c186052ca441a68ab7e5bcb6ce2ab384ff76b5756c3e7264dde6df4b.CharIn(slice[2]) || /* Match [CSZcszſ]. */
		!helpers.StartsWithIgnoreCase(slice[3:], []rune("ember")) /* Match the string "ember" (case-insensitive) */ {
		goto AlternationBranch68
	}

	alternation_branch1 = 10
	pos += 8
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch68:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 11
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("jan")) /* Match the string "jan" (case-insensitive) */ {
		goto AlternationBranch69
	}

	alternation_branch1 = 11
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch69:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 12
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("feb")) /* Match the string "feb" (case-insensitive) */ {
		goto AlternationBranch70
	}

	alternation_branch1 = 12
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch70:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 13
	// Node: Concatenate
	if len(slice) < 3 ||
		(slice[0]|0x20 != 'm') || /* Match [Mm]. */
		!set_0aed9828638bbe7e52933e044a08751d611ecbcf65fa3d7b353e7a1eecc9c411.CharIn(slice[1]) || /* Match [AaÄä]. */
		(slice[2]|0x20 != 'r') /* Match [Rr]. */ {
		goto AlternationBranch71
	}

	alternation_branch1 = 13
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch71:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 14
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("apr")) /* Match the string "apr" (case-insensitive) */ {
		goto AlternationBranch72
	}

	alternation_branch1 = 14
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch72:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 15
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("ju")) || /* Match the string "ju" (case-insensitive) */
		!helpers.IsInMask64(slice[2]-'L', 0xa0000000a0000000) /* Match [LNln]. */ {
		goto AlternationBranch73
	}

	alternation_branch1 = 15
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch73:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 16
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("aug")) /* Match the string "aug" (case-insensitive) */ {
		goto AlternationBranch74
	}

	alternation_branch1 = 16
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch74:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 17
	// Node: Concatenate
	if len(slice) < 3 ||
		((slice[0]|0x20 != 's') && (slice[0] != 'ſ')) || /* Match [Ssſ]. */
		!helpers.StartsWithIgnoreCase(slice[1:], []rune("ep")) /* Match the string "ep" (case-insensitive) */ {
		goto AlternationBranch75
	}

	alternation_branch1 = 17
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch75:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 18
	// Node: Concatenate
	if len(slice) < 3 ||
		(slice[0]|0x20 != 'o') || /* Match [Oo]. */
		!set_c585570961bd1835f850d6a9b56e3766bcb037a62b1dbc82686b27dd542a929b.CharIn(slice[1]) || /* Match [CKckK]. */
		(slice[2]|0x20 != 't') /* Match [Tt]. */ {
		goto AlternationBranch76
	}

	alternation_branch1 = 18
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch76:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 19
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("nov")) /* Match the string "nov" (case-insensitive) */ {
		goto AlternationBranch77
	}

	alternation_branch1 = 19
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch77:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 20
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("de")) || /* Match the string "de" (case-insensitive) */
		!helpers.IsInMask64(slice[2]-'C', 0x8000010080000100) /* Match [CZcz]. */ {
		goto AlternationBranch78
	}

	alternation_branch1 = 20
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch78:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 21
	// Node: Concatenate
	if len(slice) < 7 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("januari")) /* Match the string "januari" (case-insensitive) */ {
		goto AlternationBranch79
	}

	alternation_branch1 = 21
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch79:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 22
	// Node: Concatenate
	if len(slice) < 8 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("februari")) /* Match the string "februari" (case-insensitive) */ {
		goto AlternationBranch80
	}

	alternation_branch1 = 22
	pos += 8
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch80:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 23
	// Node: Concatenate
	// Node: Set(Set = [Mm])
	// Match [Mm].
	if len(slice) == 0 || (slice[0]|0x20 != 'm') {
		goto AlternationBranch81
	}

	// Node: Alternate
	// Match with 2 alternative expressions.
	if len(slice) < 2 {
		goto AlternationBranch81
	}

	switch slice[1] {
	case 'A', 'a':
		// Node: Concatenate
		// Node: Empty

		if len(slice) < 5 ||
			!helpers.StartsWithIgnoreCase(slice[2:], []rune("ret")) /* Match the string "ret" (case-insensitive) */ {
			goto AlternationBranch81
		}

		pos += 5
		slice = r.Runtext[pos:]

	case 'E', 'e':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Ii])
		// Match [Ii].
		if len(slice) < 3 || (slice[2]|0x20 != 'i') {
			goto AlternationBranch81
		}

		pos += 3
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch81
	}

	alternation_branch1 = 23
	goto AlternationMatch2

AlternationBranch81:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 24
	// Node: Concatenate
	if len(slice) < 7 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("agu")) || /* Match the string "agu" (case-insensitive) */
		((slice[3]|0x20 != 's') && (slice[3] != 'ſ')) || /* Match [Ssſ]. */
		!helpers.StartsWithIgnoreCase(slice[4:], []rune("tu")) || /* Match the string "tu" (case-insensitive) */
		((slice[6]|0x20 != 's') && (slice[6] != 'ſ')) /* Match [Ssſ]. */ {
		goto AlternationBranch82
	}

	alternation_branch1 = 24
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch82:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 25
	// Node: Concatenate
	if len(slice) < 6 ||
		(slice[0]|0x20 != 'j') || /* Match [Jj]. */
		(slice[1]|0x20 != 'ä') || /* Match [Ää]. */
		!helpers.StartsWithIgnoreCase(slice[2:], []rune("nner")) /* Match the string "nner" (case-insensitive) */ {
		goto AlternationBranch83
	}

	alternation_branch1 = 25
	pos += 6
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch83:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 26
	// Node: Concatenate
	if len(slice) < 5 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("feber")) /* Match the string "feber" (case-insensitive) */ {
		goto AlternationBranch84
	}

	alternation_branch1 = 26
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch84:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 27
	// Node: Concatenate
	if len(slice) < 4 ||
		(slice[0]|0x20 != 'm') || /* Match [Mm]. */
		(slice[1]|0x20 != 'ä') || /* Match [Ää]. */
		!helpers.StartsWithIgnoreCase(slice[2:], []rune("rz")) /* Match the string "rz" (case-insensitive) */ {
		goto AlternationBranch85
	}

	alternation_branch1 = 27
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch85:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 28
	// Node: Concatenate
	if len(slice) < 7 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("janvier")) /* Match the string "janvier" (case-insensitive) */ {
		goto AlternationBranch86
	}

	alternation_branch1 = 28
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch86:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 29
	// Node: Concatenate
	if len(slice) < 7 ||
		(slice[0]|0x20 != 'f') || /* Match [Ff]. */
		(slice[1]|0x20 != 'é') || /* Match [Éé]. */
		!helpers.StartsWithIgnoreCase(slice[2:], []rune("vrier")) /* Match the string "vrier" (case-insensitive) */ {
		goto AlternationBranch87
	}

	alternation_branch1 = 29
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch87:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 30
	// Node: Concatenate
	if len(slice) < 4 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("mar")) || /* Match the string "mar" (case-insensitive) */
		((slice[3]|0x20 != 's') && (slice[3] != 'ſ')) /* Match [Ssſ]. */ {
		goto AlternationBranch88
	}

	alternation_branch1 = 30
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch88:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 31
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("jui")) /* Match the string "jui" (case-insensitive) */ {
		goto AlternationBranch89
	}

	// Node: Alternate
	// Match with 2 alternative expressions.
	if len(slice) < 4 {
		goto AlternationBranch89
	}

	switch slice[3] {
	case 'N', 'n':
		pos += 4
		slice = r.Runtext[pos:]

	case 'L', 'l':
		// Node: Concatenate
		if len(slice) < 7 ||
			!helpers.StartsWithIgnoreCase(slice[3:], []rune("llet")) /* Match the string "llet" (case-insensitive) */ {
			goto AlternationBranch89
		}

		pos += 7
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch89
	}

	alternation_branch1 = 31
	goto AlternationMatch2

AlternationBranch89:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 32
	// Node: Concatenate
	if len(slice) < 4 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("aout")) /* Match the string "aout" (case-insensitive) */ {
		goto AlternationBranch90
	}

	alternation_branch1 = 32
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch90:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 33
	// Node: Concatenate
	if len(slice) < 9 ||
		((slice[0]|0x20 != 's') && (slice[0] != 'ſ')) || /* Match [Ssſ]. */
		!helpers.StartsWithIgnoreCase(slice[1:], []rune("eptembre")) /* Match the string "eptembre" (case-insensitive) */ {
		goto AlternationBranch91
	}

	alternation_branch1 = 33
	pos += 9
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch91:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 34
	// Node: Concatenate
	if len(slice) < 7 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("octobre")) /* Match the string "octobre" (case-insensitive) */ {
		goto AlternationBranch92
	}

	alternation_branch1 = 34
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch92:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 35
	// Node: Concatenate
	if len(slice) < 8 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("novembre")) /* Match the string "novembre" (case-insensitive) */ {
		goto AlternationBranch93
	}

	alternation_branch1 = 35
	pos += 8
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch93:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 36
	// Node: Concatenate
	if len(slice) < 8 ||
		(slice[0]|0x20 != 'd') || /* Match [Dd]. */
		(slice[1]|0x20 != 'é') || /* Match [Éé]. */
		!helpers.StartsWithIgnoreCase(slice[2:], []rune("cembre")) /* Match the string "cembre" (case-insensitive) */ {
		goto AlternationBranch94
	}

	alternation_branch1 = 36
	pos += 8
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch94:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 37
	// Node: Concatenate
	if len(slice) < 4 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("oca")) || /* Match the string "oca" (case-insensitive) */
		((slice[3]|0x20 != 'k') && (slice[3] != 'K')) /* Match [KkK]. */ {
		goto AlternationBranch95
	}

	alternation_branch1 = 37
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch95:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 38
	// Node: Concatenate
	if len(slice) < 5 ||
		!helpers.IsBetween(slice[0], 'Ş', 'ş') || /* Match [Şş]. */
		!helpers.StartsWithIgnoreCase(slice[1:], []rune("ubat")) /* Match the string "ubat" (case-insensitive) */ {
		goto AlternationBranch96
	}

	alternation_branch1 = 38
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch96:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 39
	// Node: Concatenate
	if len(slice) < 4 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("mart")) /* Match the string "mart" (case-insensitive) */ {
		goto AlternationBranch97
	}

	alternation_branch1 = 39
	pos += 4
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch97:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 40
	// Node: Concatenate
	if len(slice) < 5 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("ni")) || /* Match the string "ni" (case-insensitive) */
		((slice[2]|0x20 != 's') && (slice[2] != 'ſ')) || /* Match [Ssſ]. */
		!helpers.StartsWithIgnoreCase(slice[3:], []rune("an")) /* Match the string "an" (case-insensitive) */ {
		goto AlternationBranch98
	}

	alternation_branch1 = 40
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch98:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 41
	// Node: Concatenate
	if len(slice) < 5 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("may")) || /* Match the string "may" (case-insensitive) */
		slice[3] != 'ı' || /* Match 'ı'. */
		((slice[4]|0x20 != 's') && (slice[4] != 'ſ')) /* Match [Ssſ]. */ {
		goto AlternationBranch99
	}

	alternation_branch1 = 41
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch99:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 42
	// Node: Concatenate
	if len(slice) < 7 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("haziran")) /* Match the string "haziran" (case-insensitive) */ {
		goto AlternationBranch100
	}

	alternation_branch1 = 42
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch100:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 43
	// Node: Concatenate
	if len(slice) < 6 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("temmuz")) /* Match the string "temmuz" (case-insensitive) */ {
		goto AlternationBranch101
	}

	alternation_branch1 = 43
	pos += 6
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch101:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 44
	// Node: Concatenate
	if len(slice) < 7 ||
		(slice[0]|0x20 != 'a') || /* Match [Aa]. */
		!helpers.IsBetween(slice[1], 'Ğ', 'ğ') || /* Match [Ğğ]. */
		(slice[2]|0x20 != 'u') || /* Match [Uu]. */
		((slice[3]|0x20 != 's') && (slice[3] != 'ſ')) || /* Match [Ssſ]. */
		!helpers.StartsWithIgnoreCase(slice[4:], []rune("to")) || /* Match the string "to" (case-insensitive) */
		((slice[6]|0x20 != 's') && (slice[6] != 'ſ')) /* Match [Ssſ]. */ {
		goto AlternationBranch102
	}

	alternation_branch1 = 44
	pos += 7
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch102:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 45
	// Node: Concatenate
	// Node: Set(Set = [Ee])
	// Match [Ee].
	if len(slice) == 0 || (slice[0]|0x20 != 'e') {
		goto AlternationBranch103
	}

	// Node: Alternate
	// Match with 2 alternative expressions.
	if len(slice) < 2 {
		goto AlternationBranch103
	}

	switch slice[1] {
	case 'Y', 'y':
		// Node: Concatenate
		// Node: Empty

		if len(slice) < 5 ||
			(slice[2]|0x20 != 'l') || /* Match [Ll]. */
			(slice[3]|0x20 != 'ü') || /* Match [Üü]. */
			(slice[4]|0x20 != 'l') /* Match [Ll]. */ {
			goto AlternationBranch103
		}

		pos += 5
		slice = r.Runtext[pos:]

	case 'K', 'k', 'K':
		// Node: Concatenate
		// Node: Empty

		if len(slice) < 4 ||
			!helpers.StartsWithIgnoreCase(slice[2:], []rune("im")) /* Match the string "im" (case-insensitive) */ {
			goto AlternationBranch103
		}

		pos += 4
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch103
	}

	alternation_branch1 = 45
	goto AlternationMatch2

AlternationBranch103:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 46
	// Node: Concatenate
	if len(slice) < 5 ||
		((slice[0]|0x20 != 'k') && (slice[0] != 'K')) || /* Match [KkK]. */
		(slice[1]|0x20 != 'a') || /* Match [Aa]. */
		((slice[2]|0x20 != 's') && (slice[2] != 'ſ')) || /* Match [Ssſ]. */
		slice[3] != 'ı' || /* Match 'ı'. */
		(slice[4]|0x20 != 'm') /* Match [Mm]. */ {
		goto AlternationBranch104
	}

	alternation_branch1 = 46
	pos += 5
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch104:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 47
	// Node: Concatenate
	if len(slice) < 6 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("aral")) || /* Match the string "aral" (case-insensitive) */
		slice[4] != 'ı' || /* Match 'ı'. */
		((slice[5]|0x20 != 'k') && (slice[5] != 'K')) /* Match [KkK]. */ {
		goto AlternationBranch105
	}

	alternation_branch1 = 47
	pos += 6
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch105:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 48
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("oca")) /* Match the string "oca" (case-insensitive) */ {
		goto AlternationBranch106
	}

	alternation_branch1 = 48
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch106:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 49
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.IsBetween(slice[0], 'Ş', 'ş') || /* Match [Şş]. */
		!helpers.StartsWithIgnoreCase(slice[1:], []rune("ub")) /* Match the string "ub" (case-insensitive) */ {
		goto AlternationBranch107
	}

	alternation_branch1 = 49
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch107:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 50
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("mar")) /* Match the string "mar" (case-insensitive) */ {
		goto AlternationBranch108
	}

	alternation_branch1 = 50
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch108:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 51
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("ni")) || /* Match the string "ni" (case-insensitive) */
		((slice[2]|0x20 != 's') && (slice[2] != 'ſ')) /* Match [Ssſ]. */ {
		goto AlternationBranch109
	}

	alternation_branch1 = 51
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch109:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 52
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("haz")) /* Match the string "haz" (case-insensitive) */ {
		goto AlternationBranch110
	}

	alternation_branch1 = 52
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch110:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 53
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("tem")) /* Match the string "tem" (case-insensitive) */ {
		goto AlternationBranch111
	}

	alternation_branch1 = 53
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch111:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 54
	// Node: Concatenate
	if len(slice) < 3 ||
		(slice[0]|0x20 != 'a') || /* Match [Aa]. */
		!helpers.IsBetween(slice[1], 'Ğ', 'ğ') || /* Match [Ğğ]. */
		(slice[2]|0x20 != 'u') /* Match [Uu]. */ {
		goto AlternationBranch112
	}

	alternation_branch1 = 54
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch112:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 55
	// Node: Concatenate
	// Node: Set(Set = [Ee])
	// Match [Ee].
	if len(slice) == 0 || (slice[0]|0x20 != 'e') {
		goto AlternationBranch113
	}

	// Node: Alternate
	// Match with 2 alternative expressions.
	if len(slice) < 2 {
		goto AlternationBranch113
	}

	switch slice[1] {
	case 'Y', 'y':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Ll])
		// Match [Ll].
		if len(slice) < 3 || (slice[2]|0x20 != 'l') {
			goto AlternationBranch113
		}

		pos += 3
		slice = r.Runtext[pos:]

	case 'K', 'k', 'K':
		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [Ii])
		// Match [Ii].
		if len(slice) < 3 || (slice[2]|0x20 != 'i') {
			goto AlternationBranch113
		}

		pos += 3
		slice = r.Runtext[pos:]

	default:
		goto AlternationBranch113
	}

	alternation_branch1 = 55
	goto AlternationMatch2

AlternationBranch113:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 56
	// Node: Concatenate
	if len(slice) < 3 ||
		((slice[0]|0x20 != 'k') && (slice[0] != 'K')) || /* Match [KkK]. */
		(slice[1]|0x20 != 'a') || /* Match [Aa]. */
		((slice[2]|0x20 != 's') && (slice[2] != 'ſ')) /* Match [Ssſ]. */ {
		goto AlternationBranch114
	}

	alternation_branch1 = 56
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBranch114:
	pos = alternation_starting_pos2
	slice = r.Runtext[pos:]
	r.UncaptureUntil(alternation_starting_capturepos2)

	// Branch 57
	// Node: Concatenate
	if len(slice) < 3 ||
		!helpers.StartsWithIgnoreCase(slice, []rune("ara")) /* Match the string "ara" (case-insensitive) */ {
		goto LoopIterationNoMatch2
	}

	alternation_branch1 = 57
	pos += 3
	slice = r.Runtext[pos:]
	goto AlternationMatch2

AlternationBacktrack2:
	if err := r.CheckTimeout(); err != nil {
		return err
	}
	switch alternation_branch1 {
	case 0:
		goto CharLoopBacktrack2
	case 1:
		goto CharLoopBacktrack3
	case 2:
		goto AlternationBranch60
	case 3:
		goto AlternationBranch61
	case 4:
		goto AlternationBranch62
	case 5:
		goto AlternationBranch63
	case 6:
		goto AlternationBranch64
	case 7:
		goto AlternationBranch65
	case 8:
		goto AlternationBranch66
	case 9:
		goto AlternationBranch67
	case 10:
		goto AlternationBranch68
	case 11:
		goto AlternationBranch69
	case 12:
		goto AlternationBranch70
	case 13:
		goto AlternationBranch71
	case 14:
		goto AlternationBranch72
	case 15:
		goto AlternationBranch73
	case 16:
		goto AlternationBranch74
	case 17:
		goto AlternationBranch75
	case 18:
		goto AlternationBranch76
	case 19:
		goto AlternationBranch77
	case 20:
		goto AlternationBranch78
	case 21:
		goto AlternationBranch79
	case 22:
		goto AlternationBranch80
	case 23:
		goto AlternationBranch81
	case 24:
		goto AlternationBranch82
	case 25:
		goto AlternationBranch83
	case 26:
		goto AlternationBranch84
	case 27:
		goto AlternationBranch85
	case 28:
		goto AlternationBranch86
	case 29:
		goto AlternationBranch87
	case 30:
		goto AlternationBranch88
	case 31:
		goto AlternationBranch89
	case 32:
		goto AlternationBranch90
	case 33:
		goto AlternationBranch91
	case 34:
		goto AlternationBranch92
	case 35:
		goto AlternationBranch93
	case 36:
		goto AlternationBranch94
	case 37:
		goto AlternationBranch95
	case 38:
		goto AlternationBranch96
	case 39:
		goto AlternationBranch97
	case 40:
		goto AlternationBranch98
	case 41:
		goto AlternationBranch99
	case 42:
		goto AlternationBranch100
	case 43:
		goto AlternationBranch101
	case 44:
		goto AlternationBranch102
	case 45:
		goto AlternationBranch103
	case 46:
		goto AlternationBranch104
	case 47:
		goto AlternationBranch105
	case 48:
		goto AlternationBranch106
	case 49:
		goto AlternationBranch107
	case 50:
		goto AlternationBranch108
	case 51:
		goto AlternationBranch109
	case 52:
		goto AlternationBranch110
	case 53:
		goto AlternationBranch111
	case 54:
		goto AlternationBranch112
	case 55:
		goto AlternationBranch113
	case 56:
		goto AlternationBranch114
	case 57:
		goto LoopIterationNoMatch2
	}

AlternationMatch2:
	;

	r.Capture(5, capture_starting_pos4, pos)

	goto CaptureSkipBacktrack3

CaptureBacktrack3:
	goto AlternationBacktrack2

CaptureSkipBacktrack3:
	;

	// Node: SetloopAtomic(Set = [,\.])(Min = 0, Max = 1)
	// Match [,\.] atomically, optionally.
	if len(slice) > 0 && (slice[0]|0x2 == '.') {
		slice = slice[1:]
		pos++
	}

	// Node: Set(Set = [\s])
	// Match [\s].
	if len(slice) == 0 || !unicode.IsSpace(slice[0]) {
		goto CaptureBacktrack3
	}

	// Node: Capture(index = 6, unindex = -1)
	// "6" capture group
	pos++
	slice = r.Runtext[pos:]
	capture_starting_pos5 = pos

	// Node: Atomic
	// Node: Alternate
	// Match with 2 alternative expressions, atomically.
	if len(slice) == 0 {
		goto CaptureBacktrack3
	}

	switch slice[0] {
	case '1':
		// Node: Multi(String = "99")
		// Match the string "99".
		if !helpers.StartsWith(slice[1:], []rune("99")) {
			goto CaptureBacktrack3
		}

		// Node: Concatenate
		// Node: Empty

		// Node: Set(Set = [0-9])
		// Match [0-9].
		if len(slice) < 4 || !helpers.IsBetween(slice[3], '0', '9') {
			goto CaptureBacktrack3
		}

		pos += 4
		slice = r.Runtext[pos:]

	case '2':
		// Node: One(Ch = 0)
		// Match '0'.
		if len(slice) < 2 || slice[1] != '0' {
			goto CaptureBacktrack3
		}

		// Node: Concatenate
		// Node: Empty

		if len(slice) < 4 ||
			!helpers.IsBetween(slice[2], '0', '3') || /* Match [0-3]. */
			!helpers.IsBetween(slice[3], '0', '9') /* Match [0-9]. */ {
			goto CaptureBacktrack3
		}

		pos += 4
		slice = r.Runtext[pos:]

	default:
		goto CaptureBacktrack3
	}

	r.Capture(6, capture_starting_pos5, pos)

AlternationMatch:
	;

	r.Runstackpos = atomic_stackpos

	// The input matched.
	r.Runtextpos = pos
	r.Capture(0, matchStart, pos)
	// just to prevent an unused var error in certain regex's
	var _ = slice
	return nil
}

// The set [^\#\?\s]
var set_d224ef67dc0f70cf5e774cc9c41501d7c305e2a2d786eb4c19dc0484e59b68fd = syntax.NewCharSetRuntime("\x01\x02\x00\x00\x00\x01\x00\x00\x00##??\x01 ")

// Supports searching for the chars in or not in "0123456789ADEFHJKMNOSTadefhjkmnostŞşſK"
var sNonAscii2425845632ac04d3dc6b5f60352d44d25fa35aaf888d12722b964c9ace2b4cd4 = helpers.NewRuneSearchValues("0123456789ADEFHJKMNOSTadefhjkmnostŞşſK")

// The set [CSZcszſ]
var set_54fec145c186052ca441a68ab7e5bcb6ce2ab384ff76b5756c3e7264dde6df4b = syntax.NewCharSetRuntime("\x00\a\x00\x00\x00\x00\x00\x00\x00CCSSZZccsszzſſ")

// The set [\+-\.\w]
var set_f8012346c36030544407c813dd9113ad6c1c8080037c9de30fcfa701bc6d3567 = syntax.NewCharSetRuntime("\x00\x02\x00\x00\x00\x01\x00\x00\x00++-.\x01W")

// The set [-\.\w]
var set_bc507dd5628b7f62f93ecd744fb34f2123f824a49c4534ea1e4a3bc3e53aed33 = syntax.NewCharSetRuntime("\x00\x01\x00\x00\x00\x01\x00\x00\x00-.\x01W")

// The set [^\#/\?\s]
var set_201e66754544c28afd79d8418937f5372d668f4011ef28456fbb0936ea3e0238 = syntax.NewCharSetRuntime("\x01\x03\x00\x00\x00\x01\x00\x00\x00##//??\x01 ")

// The set [^\#\s]
var set_0811f1ec56483a6ff109ad6cdaabc3ae0d433e896a4b209895d649eca8b8e591 = syntax.NewCharSetRuntime("\x01\x01\x00\x00\x00\x01\x00\x00\x00##\x01 ")

// The set [CKckK]
var set_c585570961bd1835f850d6a9b56e3766bcb037a62b1dbc82686b27dd542a929b = syntax.NewCharSetRuntime("\x00\x05\x00\x00\x00\x00\x00\x00\x00CCKKcckkKK")

// The set [AaÄä]
var set_0aed9828638bbe7e52933e044a08751d611ecbcf65fa3d7b353e7a1eecc9c411 = syntax.NewCharSetRuntime("\x00\x04\x00\x00\x00\x00\x00\x00\x00AAaaÄÄää")

func init() {
	regexp2.RegisterEngine("[\\w\\.+-]+@[\\w\\.-]+\\.[\\w\\.-]+", regexp2.None, &rxEmail_Engine{})
	regexp2.RegisterEngine("[\\w]+://[^/\\s?#]+[^\\s?#]+(?:\\?[^\\s#]*)?(?:#[^\\s]*)?", regexp2.None, &rxURI_Engine{})
	regexp2.RegisterEngine("(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])", regexp2.None, &rxIP_Engine{})
	regexp2.RegisterEngine("(?i)(January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|Maret|Mei|Agustus|Jänner|Feber|März|janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara)\\s([0-3]?[0-9])(?:st|nd|rd|th)?,?\\s(199[0-9]|20[0-3][0-9])|([0-3]?[0-9])(?:st|nd|rd|th|\\.)?\\s(?:of\\s)?(January?|February?|March|A[pv]ril|Ma[iy]|Jun[ei]|Jul[iy]|August|September|O[ck]tober|November|De[csz]ember|Jan|Feb|M[aä]r|Apr|Jun|Jul|Aug|Sep|O[ck]t|Nov|De[cz]|Januari|Februari|Maret|Mei|Agustus|Jänner|Feber|März|janvier|février|mars|juin|juillet|aout|septembre|octobre|novembre|décembre|Ocak|Şubat|Mart|Nisan|Mayıs|Haziran|Temmuz|Ağustos|Eylül|Ekim|Kasım|Aralık|Oca|Şub|Mar|Nis|Haz|Tem|Ağu|Eyl|Eki|Kas|Ara)[,.]?\\s(199[0-9]|20[0-3][0-9])", regexp2.None, &rxLongDate_Engine{})
	var _ = helpers.Min
	var _ = syntax.NewCharSetRuntime
	var _ = unicode.IsDigit
}
