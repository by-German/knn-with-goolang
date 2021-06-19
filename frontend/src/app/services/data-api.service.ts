import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Data } from "../models/data"

@Injectable({
  providedIn: 'root'
})
export class DataApiService {
  basePath = "http://localhost:8080";
  constructor(private http: HttpClient ) { }

  getAllData(): any {
    return this.http.get<Data[]>(this.basePath + "/list")
  }

  postData( k: number, data : Data) {
    data.Discapacidad = 0
    return this.http.post<Data>(this.basePath + `/knn?k=${k}`, data)
  }
}
