**BCO-DMO Use Case**

BCO-DMO will have many DataCatalogs (representing a Cruise or a funded Project) for which there will be 1-20 Datasets attached. THis number of DataCatalogs is too large to compile into a single JSON-LD document for the BCO-DMO Organization. The selected structure of our documents has children link to their parents (Dataset > DataCatalog; DataCatalog > Organization) and parents do not define their children. To build a JSON-LD frame with the format of Organization > DataCatalog > Dataset we need to walk use the JSON-LD @reverse property to assist the framing algorithm with constructing the graph.

The relationship between the node objects in the BCO-DMO data are:
```
:dataset a schemaorg:Dataset ;
  schemaorg:includedInDataCatalog :catalog .
  
:catalog a schemaorg:DataCatalog ;
  schemaorg:publisher :organization .
  
:organization a schemaorg:Organization .
```

The data are:
```
[{
  "@context": {
        "@vocab": "http://schema.org/",
        "re3data": "http://example.org/re3data/0.1/"
    },
    "@type": "Organization",
    "@id": "http://lod.bco-dmo.org/id/affiliation/191",
    "identifier": "http://lod.bco-dmo.org/id/affiliation/191",
    "name": "BCO-DMO",
    "description": "The Biological and Chemical Oceanography Data Management Office (BCO-DMO) was created in late 2006 to serve PIs funded by the NSF Geosciences Directorate (GEO) Division of Ocean Sciences (OCE) Biological and Chemical Oceanography Programs and Office of Polar Programs (OPP) Antarctic Sciences (ANT) Organisms & Ecosystems Program. The BCO-DMO is a combination the Data Management Offices formerly created to support the US JGOFS and US GLOBEC programs. The BCO-DMO staff members are the curators of the legacy data collections created by those respective programs, as well as many other more recent research efforts including those of individual investigators.",
    "re3data:keywords": "Biological Oceanography, Chemical Oceanography",
    "contactPoint": {
        "@type": "ContactPoint",
        "name": "Adam Shepherd",
        "email": "ashepherd@whoi.edu",
        "url": "http://orcid.org/0000-0003-4486-9448",
        "contactType": "technical support"
    },
    "logo": {
        "@type": "ImageObject",
        "url": "http://www.bco-dmo.org/files/bcodmo/images/bco-dmo-words-BLUE.jpg"
    },
    "url": "http://www.bco-dmo.org",
    "sameAs": "http://www.re3data.org/repository/r3d100000012",
    "funder": {
        "@type": "Organization",
        "name": "NSF",
        "url": "http://www.nsf.gov"
    },
    "memberOf": {
        "@type": "ProgramMembership",
        "programName": "EarthCube CDF Registry",
        "hostingOrganization": {
            "@type": "Organization",
            "name": "RE3Data",
            "url": "http://www.re3data.org"
        }
    },
    "potentialAction": [{
            "@type": "SearchAction",
            "target": {
                "@type": "EntryPoint",
                "urlTemplate": "http://lod.bco-dmo.org/sparql",
                "description": "SPARQL Endpoint",
                "httpMethod": "GET"
            }
        },
        {
            "@type": "SearchAction",
            "target": {
                "@type": "EntryPoint",
                "urlTemplate": "http://www.bco-dmo.org/.well-known/void",
                "description": "VoID Description",
                "httpMethod": "GET"
            }
        },
        {
            "@type": "SearchAction",
            "target": {
                "@type": "EntryPoint",
                "urlTemplate": "http://www.bco-dmo.org/services/oai",
                "description": "OAI-PMH Endpoint",
                "httpMethod": "GET"
            }
        }
    ]
},{
   "@context":"http://schema.org/",
   "@id":"http://lod.bco-dmo.org/id/project/2227",
   "identifier":"http://lod.bco-dmo.org/id/project/2227",
   "url":"http://www.bco-dmo.org/project/2227",
   "@type":"DataCatalog",
   "name":"Santa Barbara Coastal Long Term Ecological Research site",
   "publisher":{
      "@type":"Organization",
      "name":"BCO-DMO",
      "@id":"http://lod.bco-dmo.org/id/affiliation/191",
      "identifier":"http://lod.bco-dmo.org/id/affiliation/191",
      "url":"http://www.bco-dmo.org/affiliation/191",
      "sameAs":"http://www.re3data.org/repository/r3d100000012"
   },
   "license":"http://creativecommons.org/licenses/by/4.0/",
   "citation":"http://sbc.lternet.edu/",
   "alternateName":"SBC LTER",
   "image":"http://data.bco-dmo.org/images/logos/logo_SB_LTER.png",
   "description":"u003Cpu003Eu003Cstrongu003EFrom u003Ca href=u0022http://www.lternet.edu/sites/sbcu0022u003Ehttp://www.lternet.edu/sites/sbcu003C/au003Eu003C/strongu003Eu003Cbr /u003EnThe Santa Barbara Coastal LTER is located in the coastal zone of southern California near Santa Barbara. It is bounded by the steep east-west trending Santa Ynez Mountains and coastal plain to the north and the unique Northern Channel Islands archipelago to the south. Santa Barbara Coastal Long-Term Ecological Research (SBC) Project is headquartered at the University of California, Santa Barbara, and is part of the National Science Foundationu2019s (NSF) Long-Term Ecological Research (LTER) Network.u003C/pu003Enu003Cpu003EThe research focus of SBC LTER is on ecological systems at the land-ocean margin. Although there is increasing concern about the impacts of human activities on coastal watersheds and nearshore marine environments, there have been few long-term studies of the linkages among oceanic, reef, sandy beaches, wetland, and upland habitats. SBC LTER is helping to fill this gap by studying the effects of oceanic and coastal watershed influences on kelp forests in the Santa Barbara Channel located off the coast of southern California. The primary research objective of SBC LTER is to investigate the relative importance of land vs. ocean processes in structuring giant kelp (Macrocystis pyrifera) forest ecosystems for different conditions of land use, climate and ocean influences.u003C/pu003Enu003Cpu003Eu003Cstrongu003ESBC LTER Datau003C/strongu003E: The Santa Barbara Coastal (SBC) LTER data are managed by and available directly from the SBC project data site URL shown above.u00a0 If there are any datasets listed below, they are data sets that were collected at or near the SBC LTER sampling locations, and funded by NSF OCE as ancillary projects related to the SBC LTER core research themes. See the u003Ca href=u0022http://sbc.lternet.edu/data/u0022 target=u0022_blanku0022u003ESBC LTER Data Overviewu003C/au003E page for access to data and information about data management policies.u003C/pu003En",
   "contentLocation":{
      "@type":"Place",
      "name":"Southern California Coastal Zone"
   },
   "author":[
      {
         "@type":"Person",
         "name":"Dr Daniel C. Reed",
         "@id":"http://lod.bco-dmo.org/id/person/563150",
         "identifier":"http://lod.bco-dmo.org/id/person/563150",
         "description":"Lead Principal Investigator",
         "affiliation":{
            "@type":"Organization",
            "name":"University of California-Santa Barbara (UCSB-MSI)",
            "@id":"http://lod.bco-dmo.org/id/affiliation/234",
            "identifier":"http://lod.bco-dmo.org/id/affiliation/234"
         }
      },
      {
         "@type":"Person",
         "name":"Dr John Melack",
         "@id":"http://lod.bco-dmo.org/id/person/51545",
         "identifier":"http://lod.bco-dmo.org/id/person/51545",
         "description":"Co-Principal Investigator",
         "affiliation":{
            "@type":"Organization",
            "name":"University of California-Santa Barbara (UCSB)",
            "@id":"http://lod.bco-dmo.org/id/affiliation/76",
            "identifier":"http://lod.bco-dmo.org/id/affiliation/76"
         }
      },
      {
         "@type":"Person",
         "name":"Dr Sally Holbrook",
         "@id":"http://lod.bco-dmo.org/id/person/51537",
         "identifier":"http://lod.bco-dmo.org/id/person/51537",
         "description":"Co-Principal Investigator",
         "affiliation":{
            "@type":"Organization",
            "name":"University of California-Santa Barbara (UCSB)",
            "@id":"http://lod.bco-dmo.org/id/affiliation/76",
            "identifier":"http://lod.bco-dmo.org/id/affiliation/76"
         }
      },
      {
         "@type":"Person",
         "name":"Dr David Siegel",
         "@id":"http://lod.bco-dmo.org/id/person/50849",
         "identifier":"http://lod.bco-dmo.org/id/person/50849",
         "description":"Co-Principal Investigator",
         "affiliation":{
            "@type":"Organization",
            "name":"University of California-Santa Barbara (UCSB-ICESS)",
            "@id":"http://lod.bco-dmo.org/id/affiliation/254",
            "identifier":"http://lod.bco-dmo.org/id/affiliation/254"
         }
      },
      {
         "@type":"Person",
         "name":"Margaret Ou0026#039;Brien",
         "@id":"http://lod.bco-dmo.org/id/person/551032",
         "identifier":"http://lod.bco-dmo.org/id/person/551032",
         "description":"Data Manager",
         "affiliation":{
            "@type":"Organization",
            "name":"University of California-Santa Barbara (UCSB-MSI)",
            "@id":"http://lod.bco-dmo.org/id/affiliation/234",
            "identifier":"http://lod.bco-dmo.org/id/affiliation/234"
         }
      }
   ]
},{
   "@context":"http://schema.org/",
   "@id":"http://lod.bco-dmo.org/id/dataset/517839",
   "identifier":"http://lod.bco-dmo.org/id/dataset/517839",
   "url":"http://www.bco-dmo.org/dataset/517839",
   "@type":"Dataset",
   "name":"Experimental and survey biogeochemical and microbial data from R/V Point Sur cruises PS1009 and PS1103 in the Santa Barbara Channel from 2010-2011 (SBDOM project, SBC LTER project)",
   "alternateName":"Exp and Survey Biogeochem and Microbial data",
   "description":"Experimental and survey biogeochemical and microbial data from SBDOM cruises in the Santa Barbara Channel.",
   "version":"08 July 2014",
   "includedInDataCatalog":[
      "http://lod.bco-dmo.org/id/deployment/58833",
      "http://lod.bco-dmo.org/id/deployment/517703",
      "http://lod.bco-dmo.org/id/project/2226",
      "http://lod.bco-dmo.org/id/project/2227"
   ]
}
  ]
```

My guess at the JSON-LD Frame is currently below (similar to EXAMPLE 2 from [JSON-LD Framing](https://json-ld.org/spec/latest/json-ld-framing/#framing)), but this doesn't work correctly. It grabs all Organizations, but doesn't set the publishesDataCatalog:
```
{
  "@context": {
    "@vocab": "http://schema.org/",
    "publishesDataCatalog": {
      "@reverse": "publisher"
    },
    "publisher": {
      "@type": "@id"
    },
    "dataset": {
       "@reverse": "includedInDataCatalog"
    },
    "includedInDataCatalog": {
      "@type": "@id",
      "@container": "@set"
    }
  },
  "@type": "Organization",
  "publishesDataCatalog": {
    "@type": "DataCatalog",
    "dataset": {
      "@type": "Dataset"
    }
  }
}
```
To confirm, I tried to load just the DataCatalog and their Datasets, but this only loads the DataCatalogs:
```
{
  "@context": {
    "@vocab": "http://schema.org/",
    "publishesDataCatalog": {
      "@reverse": "publisher"
    },
    "publisher": {
      "@type": "@id"
    },
    "dataset": {
       "@reverse": "includedInDataCatalog"
    },
    "includedInDataCatalog": {
      "@type": "@id",
      "@container": "@set"
    }
  },
  "@type": "DataCatalog",
  "dataset": {
    "@type": "Dataset"
  }  
}
```
Finally, just trying to go from Organization > DataCatalog doesn't work:
```
{
  "@context": {
    "@vocab": "http://schema.org/",
    "catalog": {
      "@reverse": "publisher"
    },
    "publisher": {
      "@type": "@id"
    }
  },
  "@type": "Organization",
  "catalog": {
    "@type": "DataCatalog"
  }  
}
```
