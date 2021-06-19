import { Component } from '@angular/core';
import { FormGroup, FormControl } from "@angular/forms";
import {DataApiService} from "../../services/data-api.service";
import {Data} from "../../models/data";
import {isObject} from "rxjs/internal-compatibility";

@Component({
  selector: 'app-form',
  templateUrl: './form.component.html',
  styleUrls: ['./form.component.css']
})
export class FormComponent {
  form = new FormGroup({
    departamento: new FormControl(''),
    parentesco: new FormControl(''),
    miembroHogar: new FormControl(''),
    edad: new FormControl(''),
    nivel: new FormControl(''),
    k: new FormControl('')
  })

  result : boolean = false
  data: Data = new class implements Data {
    Departamento: string = "";
    Discapacidad: number = 0;
    Edad: number = 0;
    MiembroHogar: number = 0;
    NivelEstudios: number = 0;
    Parentesco: number = 0;
  }

  constructor(private api: DataApiService) { }

  public showData() {
    this.data.Departamento = `${this.form.controls.departamento.value}`;
    this.data.Parentesco = this.form.controls.parentesco.value;
    this.data.MiembroHogar = this.form.controls.miembroHogar.value;
    this.data.Edad = this.form.controls.edad.value;
    this.data.NivelEstudios = this.form.controls.nivel.value;
    this.data.Discapacidad = 0;

    this.api.postData(this.form.controls.k.value, this.data)
      .subscribe( (data: Data) => {
        this.data = data
        this.result = true
      });
    return
  }

  existResult() : boolean {
    return this.result;
  }
}
