package gotest

import (
	"context"
	"errors"
	"fmt"

	commonpb "go.viam.com/api/common/v1"
	arm "go.viam.com/rdk/components/arm"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/referenceframe"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/spatialmath"
)

var (
	Arm              = resource.NewModel("allisonorg", "gotest", "arm")
	errUnimplemented = errors.New("unimplemented")
)

func init() {
	resource.RegisterComponent(arm.API, Arm,
		resource.Registration[arm.Arm, *Config]{
			Constructor: newGotestArm,
		},
	)
}

type Config struct {
	/*
		Put config attributes here. There should be public/exported fields
		with a `json` parameter at the end of each attribute.

		Example config struct:
			type Config struct {
				Pin   string `json:"pin"`
				Board string `json:"board"`
				MinDeg *float64 `json:"min_angle_deg,omitempty"`
			}

		If your model does not need a config, replace *Config in the init
		function with resource.NoNativeConfig
	*/
}

// Validate ensures all parts of the config are valid and important fields exist.
// Returns three values:
//  1. Required dependencies: other resources that must exist for this resource to work.
//  2. Optional dependencies: other resources that may exist but are not required.
//  3. An error if any Config fields are missing or invalid.
//
// The `path` parameter indicates
// where this resource appears in the machine's JSON configuration
// (for example, "components.0"). You can use it in error messages
// to indicate which resource has a problem.
func (cfg *Config) Validate(path string) ([]string, []string, error) {
	// Add config validation code here
	return nil, nil, nil
}

type gotestArm struct {
	resource.AlwaysRebuild

	name resource.Name

	logger logging.Logger
	cfg    *Config

	cancelCtx  context.Context
	cancelFunc func()
}

func newGotestArm(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (arm.Arm, error) {
	conf, err := resource.NativeConfig[*Config](rawConf)
	if err != nil {
		return nil, err
	}

	return NewArm(ctx, deps, rawConf.ResourceName(), conf, logger)

}

func NewArm(ctx context.Context, deps resource.Dependencies, name resource.Name, conf *Config, logger logging.Logger) (arm.Arm, error) {

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	s := &gotestArm{
		name:       name,
		logger:     logger,
		cfg:        conf,
		cancelCtx:  cancelCtx,
		cancelFunc: cancelFunc,
	}
	return s, nil
}

func (s *gotestArm) Name() resource.Name {
	return s.name
}

// EndPosition returns the current position of the arm.
func (s *gotestArm) EndPosition(ctx context.Context, extra map[string]interface{}) (spatialmath.Pose, error) {
	var poseRetVal spatialmath.Pose

	return poseRetVal, fmt.Errorf("not implemented")
}

// MoveToPosition moves the arm to the given absolute position.
// This will block until done or a new operation cancels this one.
func (s *gotestArm) MoveToPosition(ctx context.Context, pose spatialmath.Pose, extra map[string]interface{}) error {
	return fmt.Errorf("not implemented")
}

// MoveToJointPositions moves the arm's joints to the given positions.
// This will block until done or a new operation cancels this one.
func (s *gotestArm) MoveToJointPositions(ctx context.Context, positions []referenceframe.Input, extra map[string]interface{}) error {
	return fmt.Errorf("not implemented")
}

// MoveThroughJointPositions moves the arm's joints through the given positions in the order they are specified.
// This will block until done or a new operation cancels this one.
func (s *gotestArm) MoveThroughJointPositions(ctx context.Context, positions [][]referenceframe.Input, options *arm.MoveOptions, extra map[string]interface{}) error {
	return fmt.Errorf("not implemented")
}

// JointPositions returns the current joint positions of the arm.
func (s *gotestArm) JointPositions(ctx context.Context, extra map[string]interface{}) ([]referenceframe.Input, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *gotestArm) Stop(ctx context.Context, extra map[string]interface{}) error {
	return fmt.Errorf("not implemented")
}

func (s *gotestArm) Kinematics(ctx context.Context) (referenceframe.Model, error) {
	var modelRetVal referenceframe.Model

	return modelRetVal, fmt.Errorf("not implemented")
}

func (s *gotestArm) CurrentInputs(ctx context.Context) ([]referenceframe.Input, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *gotestArm) GoToInputs(ctx context.Context, inputSteps ...[]referenceframe.Input) error {
	return fmt.Errorf("not implemented")
}

func (s *gotestArm) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *gotestArm) IsMoving(ctx context.Context) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func (s *gotestArm) Geometries(ctx context.Context, extra map[string]interface{}) ([]spatialmath.Geometry, error) {
	return nil, fmt.Errorf("not implemented")
}

// Get3DModels returns the 3D models of the arm.
func (s *gotestArm) Get3DModels(ctx context.Context, extra map[string]interface{}) (map[string]*commonpb.Mesh, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *gotestArm) Close(context.Context) error {
	// Put close code here
	s.cancelFunc()
	return nil
}
