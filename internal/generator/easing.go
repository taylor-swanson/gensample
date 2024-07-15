package generator

import (
	"fmt"
	"strings"

	"github.com/fogleman/ease"
)

type Easing int

const (
	EasingLinear Easing = iota
	EasingInQuad
	EasingOutQuad
	EasingInOutQuad
	EasingInCubic
	EasingOutCubic
	EasingInOutCubic
	EasingInQuart
	EasingOutQuart
	EasingInOutQuart
	EasingInQuint
	EasingOutQuint
	EasingInOutQuint
	EasingInSine
	EasingOutSine
	EasingInOutSine
	EasingInExp
	EasingOutExp
	EasingInOutExp
	EasingInCirc
	EasingOutCirc
	EasingInOutCirc
)

func (e *Easing) String() string {
	switch *e {
	case EasingLinear:
		return "linear"
	case EasingInQuad:
		return "in_quad"
	case EasingOutQuad:
		return "out_quad"
	case EasingInOutQuad:
		return "in_out_quad"
	case EasingInCubic:
		return "in_cubic"
	case EasingOutCubic:
		return "out_cubic"
	case EasingInOutCubic:
		return "in_out_cubic"
	case EasingInQuart:
		return "in_quart"
	case EasingOutQuart:
		return "out_quart"
	case EasingInOutQuart:
		return "in_out_quart"
	case EasingInQuint:
		return "in_quint"
	case EasingOutQuint:
		return "out_quint"
	case EasingInOutQuint:
		return "in_out_quint"
	case EasingInSine:
		return "in_sine"
	case EasingOutSine:
		return "out_sine"
	case EasingInOutSine:
		return "in_out_sine"
	case EasingInExp:
		return "in_exp"
	case EasingOutExp:
		return "out_exp"
	case EasingInOutExp:
		return "in_out_exp"
	case EasingInCirc:
		return "in_circ"
	case EasingOutCirc:
		return "out_circ"
	case EasingInOutCirc:
		return "in_out_circ"
	}

	return ""
}

func (e *Easing) Unpack(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("easing value must be a string")
	}
	switch strings.ToLower(s) {
	case "linear":
		*e = EasingLinear
	case "in_quad":
		*e = EasingInQuad
	case "out_quad":
		*e = EasingOutQuad
	case "in_out_quad":
		*e = EasingInOutQuad
	case "in_cubic":
		*e = EasingInCubic
	case "out_cubic":
		*e = EasingOutCubic
	case "in_out_cubic":
		*e = EasingInOutCubic
	case "in_quart":
		*e = EasingInQuart
	case "out_quart":
		*e = EasingOutQuart
	case "in_out_quart":
		*e = EasingInOutQuart
	case "in_quint":
		*e = EasingInQuint
	case "out_quint":
		*e = EasingOutQuint
	case "in_out_quint":
		*e = EasingInOutQuint
	case "in_sine":
		*e = EasingInSine
	case "out_sine":
		*e = EasingOutSine
	case "in_out_sine":
		*e = EasingInOutSine
	case "in_exp":
		*e = EasingInExp
	case "out_exp":
		*e = EasingOutExp
	case "in_out_exp":
		*e = EasingInOutExp
	case "in_circ":
		*e = EasingInCirc
	case "out_circ":
		*e = EasingOutCirc
	case "in_out_circ":
		*e = EasingInOutCirc
	default:
		return fmt.Errorf("unknown easing value: %q", s)
	}

	return nil
}

func (e *Easing) Value(t float64) float64 {
	switch *e {
	case EasingLinear:
		return t
	case EasingInQuad:
		return ease.InQuad(t)
	case EasingOutQuad:
		return ease.OutQuad(t)
	case EasingInOutQuad:
		return ease.InOutQuad(t)
	case EasingInCubic:
		return ease.InCubic(t)
	case EasingOutCubic:
		return ease.OutCubic(t)
	case EasingInOutCubic:
		return ease.InOutCubic(t)
	case EasingInQuart:
		return ease.InQuart(t)
	case EasingOutQuart:
		return ease.OutQuart(t)
	case EasingInOutQuart:
		return ease.InOutQuart(t)
	case EasingInQuint:
		return ease.InQuint(t)
	case EasingOutQuint:
		return ease.OutQuint(t)
	case EasingInOutQuint:
		return ease.InOutQuint(t)
	case EasingInSine:
		return ease.InSine(t)
	case EasingOutSine:
		return ease.OutSine(t)
	case EasingInOutSine:
		return ease.InOutSine(t)
	case EasingInExp:
		return ease.InExpo(t)
	case EasingOutExp:
		return ease.OutExpo(t)
	case EasingInOutExp:
		return ease.InOutExpo(t)
	case EasingInCirc:
		return ease.InCirc(t)
	case EasingOutCirc:
		return ease.OutCirc(t)
	case EasingInOutCirc:
		return ease.InOutCirc(t)
	}

	return 0
}
