
/*----------------------------------------------------------------------------*/
/* Copyright (c) 2019 FIRST. All Rights Reserved.                             */
/* Open Source Software - may be modified and shared by FRC teams. The code   */
/* must be accompanied by the FIRST BSD license file in the root directory of */
/* the project.                                                               */
/*----------------------------------------------------------------------------*/
//TODO Finish
package {{.Package}};

import edu.wpi.first.wpilibj2.command.SubsystemBase;

import frc.robot.FRCLib.Motors.FRCVictorSPX;
import frc.robot.FRCLib.Motors.FRCTalonSRX;
import frc.robot.FRCLib.Motors.FRCTalonFX;

/**
 * {{.Class}} is a generated subsystem
 * {{.Subsystem.Summary}}
 *
 * {{.Subsystem.Description}}
 */
public class {{.Class}} extends SubsystemBase {

    {{range .Motors}}
  public {{.MotorControllerClassName}} {{.Name}};
  {{end}}

/*
  public static enum ActionState {
    INTAKING, NOT_INTAKING
  }

  public ActionState actionState;
*/
  /**
   * Creates a {{.Class}}
   * {{.Subsystem.Summary}}
   */
  public {{.Class}}() {
{{range .Motors}}
    {{.Name}} = new {{.MotorControllerClassName}}.{{.MotorControllerClassName}}Builder({{.CanID}})
        {{if .EnableInvert}}.withInverted({{.InvertEnum}}){{end}}
        {{if .EnableNeutral}}.withNeutralMode({{.NeutralEnum}}}}){{end}}
        {{if .EnableFeedback}}.withFeedbackPort({{.FeedbackPortEnum}}){{end}}
        {{if .EnableSensorPhase}}.withSensorPhase({{.SensorPhaseEnum}}){{end}}
        {{if .TimeoutEnabled}}.withTimeout({{.Timeout}}){{end}}
        .withCurrentLimitEnabled({{.EnableCurrentLimit}}) {{/*Not conditional because conditional is handled in FRCLib*/}}
        {{if .EnableCurrentLimit}}.withCurrentLimit({{.CurrentLimit}}){{end}}
        {{if .EnableOpenLoopRamp}}.withOpenLoopRampRate({{.OpenLoopRamp}}){{end}}
        {{if .EnableNominalOutput}}.withNominalOutputForward({{.NominalOutputForward}}){{end}}
        {{if .EnableNominalOutput}}.withNominalOutputReverse({{.NominalOutputReverse}}){{end}}
        {{if .EnablePeakOutput}}.withPeakOutputForward({{.PeakOutputForward}}){{end}}
        {{if .EnablePeakOutput}}.withPeakOutputReverse({{.PeakOutputReverse}}){{end}}
        {{if .EnableFollow}}.withMaster({{.Follows}}){{end}}
        .build();

        addChild("{{.Name}}", {{.Name}});
  }
  {{end}}


  /**
   * Update any states
   */
  public void updateState() {
    {{.Subsystem.OnStateUpdateCode}}
  }

  @Override
  public void periodic() {
    // This method will be called once per scheduler run
    updateState();
  }
}