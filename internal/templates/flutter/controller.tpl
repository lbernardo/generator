import 'package:mobx/mobx.dart';

part '{{.name}}_controller.g.dart';

class {{.ClassName}}Controller = {{.ClassName}}ControllerBase with _${{.ClassName}}Controller;

abstract class {{.ClassName}}ControllerBase with Store {
  {{.ClassName}}ControllerBase();
}
