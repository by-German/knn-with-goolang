import { Component, OnInit } from '@angular/core';
import { DataApiService } from "../../services/data-api.service";
import { Data } from "../../models/data"

export interface PeriodicElement {
  name: string;
  position: number;
  weight: number;
  symbol: string;
}

@Component({
  selector: 'app-screen',
  templateUrl: './screen.component.html',
  styleUrls: ['./screen.component.css']
})
export class ScreenComponent implements OnInit {
  elements : Data[] = [];

  constructor(private data : DataApiService) { }

  ngOnInit(): void {
    this.data.getAllData()
      .subscribe( (data: Data[]) => {
        this.elements = data
      });
  }

  test() {
    this.elements.pop()
    console.log(this.elements)
  }

  displayedColumns: string[] = ['Departamento', 'Parentesco', 'MiembroHogar', 'Edad', "NivelEstudios", "Discapacidad"];
}
